// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package builtin contains functions for evaluating built-in predicates.
package builtin

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/mangle/ast"
	"github.com/google/mangle/functional"
	"github.com/google/mangle/symbols"
	"github.com/google/mangle/unionfind"
)

var (
	// Predicates built-in predicates.
	Predicates = map[ast.PredicateSym]struct{}{
		symbols.Lt:             {},
		symbols.Le:             {},
		symbols.ListMember:     {},
		symbols.WithinDistance: {},
		symbols.MatchPair:      {},
		symbols.MatchCons:      {},
		symbols.MatchNil:       {},
		symbols.MatchField:     {},
		symbols.MatchEntry:     {},
	}

	// Functions has all built-in functions except reducers.
	Functions = map[ast.FunctionSym]struct{}{
		symbols.Div:   {},
		symbols.Mult:  {},
		symbols.Plus:  {},
		symbols.Minus: {},

		// This is only used to start a "do-transform".
		symbols.GroupBy: {},

		symbols.ListGet: {},
		symbols.Append:  {},
		symbols.Cons:    {},
		symbols.Len:     {},
		symbols.List:    {},
		symbols.Pair:    {},
		symbols.Tuple:   {},
	}

	// ReducerFunctions has those built-in functions with are reducers.
	ReducerFunctions = map[ast.FunctionSym]struct{}{
		symbols.Collect:         {},
		symbols.CollectDistinct: {},
		symbols.PickAny:         {},
		symbols.Max:             {},
		symbols.Sum:             {},
		symbols.Count:           {},
	}

	// errFound is used for exiting a loop
	errFound = errors.New("found")
)

func init() {
	for fn := range ReducerFunctions {
		Functions[fn] = struct{}{}
	}
}

// IsBuiltinFunction returns true if sym is a builtin function.
func IsBuiltinFunction(sym ast.FunctionSym) bool {
	if _, ok := Functions[sym]; ok {
		return true
	}
	if _, ok := Functions[ast.FunctionSym{sym.Symbol, -1}]; ok {
		return true // variable arity
	}
	return false
}

// IsReducerFunction returns true if sym is a reducer function.
func IsReducerFunction(sym ast.FunctionSym) bool {
	if _, ok := ReducerFunctions[sym]; ok {
		return true
	}
	if _, ok := ReducerFunctions[ast.FunctionSym{sym.Symbol, -1}]; ok {
		return true // variable arity
	}
	return false
}

// Decide evaluates an atom of a built-in predicate. The atom must no longer contain any
// apply-expressions or variables.
func Decide(atom ast.Atom, subst *unionfind.UnionFind) (bool, []*unionfind.UnionFind, error) {
	switch atom.Predicate.Symbol {
	case symbols.MatchPair.Symbol:
		fallthrough
	case symbols.MatchCons.Symbol:
		fallthrough
	case symbols.MatchEntry.Symbol:
		fallthrough
	case symbols.MatchField.Symbol:
		fallthrough
	case symbols.MatchNil.Symbol:
		ok, nsubst, err := match(atom, subst)
		if err != nil {
			return false, nil, err
		}
		if !ok {
			return false, nil, nil
		}
		return ok, []*unionfind.UnionFind{nsubst}, nil
	}
	switch atom.Predicate.Symbol {
	case symbols.Lt.Symbol:
		if len(atom.Args) != 2 {
			return false, nil, fmt.Errorf("wrong number of arguments for built-in predicate '<': %v", atom.Args)
		}
		nums, err := getNumberValues(atom.Args)
		if err != nil {
			return false, nil, err
		}
		return nums[0] < nums[1], []*unionfind.UnionFind{subst}, nil
	case symbols.Le.Symbol:
		if len(atom.Args) != 2 {
			return false, nil, fmt.Errorf("wrong number of arguments for built-in predicate '<=': %v", atom.Args)
		}
		nums, err := getNumberValues(atom.Args)
		if err != nil {
			return false, nil, err
		}
		return nums[0] <= nums[1], []*unionfind.UnionFind{subst}, nil

	case symbols.ListMember.Symbol: // :list:member(Member, List)
		evaluatedArg, err := functional.EvalExpr(atom.Args[1], subst)
		if err != nil {
			return false, nil, err
		}
		c, ok := evaluatedArg.(ast.Constant)
		if !ok {
			return false, nil, fmt.Errorf("not a constant: %v %T", evaluatedArg, evaluatedArg)
		}
		evaluatedMember := atom.Args[0]
		memberVar, memberIsVar := evaluatedMember.(ast.Variable)
		if memberIsVar && subst != nil {
			evaluatedMember = subst.Get(memberVar)
			_, memberIsVar = evaluatedMember.(ast.Variable)
		}
		if !memberIsVar { // We are looking for a member
			res, err := functional.EvalExpr(
				ast.ApplyFn{symbols.ListContains, []ast.BaseTerm{evaluatedArg, evaluatedMember}}, nil)
			if err != nil {
				return false, nil, err
			}
			return res.Equals(ast.TrueConstant), []*unionfind.UnionFind{subst}, nil
		}
		if c.Type != ast.ListShape {
			return false, nil, nil // If expanding fails, this is not an error.
		}
		var values []ast.Constant
		c.ListValues(func(elem ast.Constant) error {
			values = append(values, elem)
			return nil
		}, func() error { return nil })
		if len(values) > 0 {
			var nsubsts []*unionfind.UnionFind
			for _, elem := range values {
				nsubst, err := unionfind.UnifyTermsExtend([]ast.BaseTerm{memberVar}, []ast.BaseTerm{elem}, *subst)
				if err != nil {
					return false, nil, err
				}
				nsubsts = append(nsubsts, &nsubst)
			}
			return true, nsubsts, nil
		}
		return false, nil, nil

	case symbols.WithinDistance.Symbol:
		if len(atom.Args) != 3 {
			return false, nil, fmt.Errorf("wrong number of arguments for built-in predicate 'within_distance': %v", atom.Args)
		}
		nums, err := getNumberValues(atom.Args)
		if err != nil {
			return false, nil, err
		}
		return abs(nums[0]-nums[1]) < nums[2], []*unionfind.UnionFind{subst}, nil
	default:
		return false, nil, fmt.Errorf("not a builtin predicate: %s", atom.Predicate.Symbol)
	}
}

func match(pattern ast.Atom, subst *unionfind.UnionFind) (bool, *unionfind.UnionFind, error) {
	evaluatedArg, err := functional.EvalExpr(pattern.Args[0], subst)
	if err != nil {
		return false, nil, err
	}
	scrutinee, ok := evaluatedArg.(ast.Constant)
	if !ok {
		return false, nil, fmt.Errorf("not a constant: %v %T", evaluatedArg, evaluatedArg)
	}
	switch pattern.Predicate.Symbol {
	case symbols.MatchPair.Symbol:
		if len(pattern.Args) != 3 {
			return false, nil, fmt.Errorf("wrong number of arguments for built-in predicate ':match_pair': %v", pattern.Args)
		}
		leftVar, leftOK := pattern.Args[1].(ast.Variable)
		rightVar, rightOk := pattern.Args[2].(ast.Variable)
		if !leftOK || !rightOk {
			return false, nil, fmt.Errorf("2nd and 3rd arguments must be variables for ':match_pair': %v", pattern)
		}

		fst, snd, err := scrutinee.PairValue()
		if err != nil {
			return false, nil, nil // failing match is not an error
		}
		// First argument is indeed a pair. Bind.
		nsubst, err := unionfind.UnifyTermsExtend([]ast.BaseTerm{leftVar, rightVar}, []ast.BaseTerm{fst, snd}, *subst)
		if err != nil {
			return false, nil, fmt.Errorf("This should never happen for %v", pattern)
		}
		return true, &nsubst, nil

	case symbols.MatchCons.Symbol:
		if len(pattern.Args) != 3 {
			return false, nil, fmt.Errorf("wrong number of arguments for built-in predicate ':match_cons': %v", pattern.Args)
		}
		leftVar, leftOK := pattern.Args[1].(ast.Variable)
		rightVar, rightOk := pattern.Args[2].(ast.Variable)
		if !leftOK || !rightOk {
			return false, nil, fmt.Errorf("2nd and 3rd arguments must be variables for ':match_cons': %v", pattern)
		}

		scrutineeList, err := getListValue(scrutinee)
		if err != nil {
			return false, nil, nil // failing match is not an error
		}
		hd, tail, err := scrutineeList.ConsValue()
		if err != nil {
			return false, nil, nil // failing match is not an error
		}
		// First argument is indeed a cons. Bind.
		nsubst, err := unionfind.UnifyTermsExtend([]ast.BaseTerm{leftVar, rightVar}, []ast.BaseTerm{hd, tail}, *subst)
		if err != nil {
			return false, nil, fmt.Errorf("This should never happen for %v", pattern)
		}
		return true, &nsubst, nil

	case symbols.MatchNil.Symbol:
		if len(pattern.Args) != 1 {
			return false, nil, fmt.Errorf("wrong number of arguments for built-in predicate ':match_nil': %v", pattern.Args)
		}
		if !scrutinee.IsListNil() {
			return false, nil, nil
		}
		return true, subst, nil

	case symbols.MatchEntry.Symbol:
		if scrutinee.Type != ast.MapShape || scrutinee.IsMapNil() {
			return false, nil, nil
		}
		patternKey, ok := pattern.Args[1].(ast.Constant)
		if !ok {
			return false, nil, fmt.Errorf("bad pattern %v", pattern) // This should not happen
		}
		patternVal := pattern.Args[2]
		var found *ast.Constant
		e, err := scrutinee.MapValues(func(key ast.Constant, val ast.Constant) error {
			if key.Equals(patternKey) {
				found = &val
				return errFound
			}
			return nil
		}, func() error { return nil })
		if e != nil {
			return false, nil, e // This should not happen
		}
		if errors.Is(err, errFound) {
			if nsubst, errUnify := unionfind.UnifyTermsExtend([]ast.BaseTerm{patternVal}, []ast.BaseTerm{*found}, *subst); errUnify == nil { // if NO error
				return true, &nsubst, nil
			}
		}
		return false, nil, nil

	case symbols.MatchField.Symbol:
		if scrutinee.Type != ast.StructShape || scrutinee.IsStructNil() {
			return false, nil, nil
		}
		patternKey, ok := pattern.Args[1].(ast.Constant)
		if !ok {
			return false, nil, fmt.Errorf("bad pattern %v", pattern) // This should not happen
		}
		patternVal := pattern.Args[2]
		var found *ast.Constant
		e, err := scrutinee.StructValues(func(key ast.Constant, val ast.Constant) error {
			if key.Equals(patternKey) {
				found = &val
				return errFound
			}
			return nil
		}, func() error { return nil })
		if e != nil {
			return false, nil, nil // This should not happen
		}

		if errors.Is(err, errFound) {
			if nsubst, errUnify := unionfind.UnifyTermsExtend([]ast.BaseTerm{patternVal}, []ast.BaseTerm{*found}, *subst); errUnify == nil { // if NO error
				return true, &nsubst, nil
			}
		}
		return false, nil, nil

	default:
		return false, nil, fmt.Errorf("unexpected case: %v", pattern.Predicate.Symbol)
	}
}

func getStringValue(baseTerm ast.BaseTerm) (string, error) {
	constant, ok := baseTerm.(ast.Constant)
	if !ok || constant.Type != ast.StringType {
		return "", fmt.Errorf("value %v (%T) is not a string", baseTerm, baseTerm)
	}
	return constant.StringValue()
}

func getNumberValue(b ast.BaseTerm) (int64, error) {
	c, ok := b.(ast.Constant)
	if !ok {
		return 0, fmt.Errorf("not a value %v (%T)", b, b)
	}
	if c.Type != ast.NumberType {
		return 0, fmt.Errorf("value %v (%v) is not a number", c, c.Type)
	}
	return c.NumberValue()
}

func getFloatValue(b ast.BaseTerm) (float64, error) {
	c, ok := b.(ast.Constant)
	if !ok {
		return 0, fmt.Errorf("not a value %v (%T)", b, b)
	}
	if c.Type != ast.Float64Type {
		return 0, fmt.Errorf("value %v (%v) is not a number", c, c.Type)
	}
	return c.Float64Value()
}

func getListValue(c ast.Constant) (ast.Constant, error) {
	if c.Type != ast.ListShape {
		return ast.Constant{}, fmt.Errorf("value %v (%v) is not a list", c, c.Type)
	}
	return c, nil
}

func getMapValue(c ast.Constant) (ast.Constant, error) {
	if c.Type != ast.MapShape {
		return ast.Constant{}, fmt.Errorf("value %v (%v) is not a map", c, c.Type)
	}
	return c, nil
}

func getStructValue(c ast.Constant) (ast.Constant, error) {
	if c.Type != ast.StructShape {
		return ast.Constant{}, fmt.Errorf("value %v (%v) is not a struct", c, c.Type)
	}
	return c, nil
}

func getNumberValues[T ast.BaseTerm](cs []T) ([]int64, error) {
	var nums []int64
	for _, c := range cs {
		num, err := getNumberValue(c)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

// Abs returns the absolute value of x.
func abs(x int64) int64 {
	if x < 0 {
		return -x // This is wrong for math.MinInt
	}
	return x
}

// TypeChecker checks the type of constant (run-time type).
type TypeChecker struct {
	decls map[ast.PredicateSym]*ast.Decl
}

// NewTypeChecker returns a new TypeChecker.
// The decls must be desugared so they only contain type bounds.
func NewTypeChecker(decls map[ast.PredicateSym]ast.Decl) (*TypeChecker, error) {
	desugaredDecls, err := symbols.CheckAndDesugar(decls)
	if err != nil {
		return nil, err
	}
	return NewTypeCheckerFromDesugared(desugaredDecls), nil
}

// NewTypeCheckerFromDesugared returns a new TypeChecker.
// The declarations must be in desugared form.
func NewTypeCheckerFromDesugared(decls map[ast.PredicateSym]*ast.Decl) *TypeChecker {
	return &TypeChecker{decls}
}

// CheckTypeBounds checks whether there fact is consistent with at least one of the bound decls.
func (t TypeChecker) CheckTypeBounds(fact ast.Atom) error {
	decl, ok := t.decls[fact.Predicate]
	if !ok {
		return fmt.Errorf("could not find declaration for %v", fact.Predicate)
	}
	var errs []string
	for _, boundDecl := range decl.Bounds {
		err := t.CheckOneBoundDecl(fact, boundDecl)
		if err == nil { // if NO error
			return nil
		}
		errs = append(errs, fmt.Sprintf("bound decl %v fails with %v", boundDecl, err))
	}
	return fmt.Errorf("fact %v matches none of the bound decls: %v", fact, strings.Join(errs, ","))
}

// CheckOneBoundDecl checks whether a fact is consistent with a given type bounds tuple.
func (t TypeChecker) CheckOneBoundDecl(fact ast.Atom, boundDecl ast.BoundDecl) error {
	for j, bound := range boundDecl.Bounds {
		c, ok := fact.Args[j].(ast.Constant)
		if !ok {
			return fmt.Errorf("fact %v could not check fact argument %v", fact, fact.Args[j])
		}
		t, err := symbols.NewTypeHandle(bound)
		if err != nil {
			return err
		}
		if !t.HasType(c) {
			return fmt.Errorf("argument %v is not an element of |%v|", c, t)
		}
	}
	return nil
}
