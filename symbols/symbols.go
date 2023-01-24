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

// Package symbols contains symbols for built-in functions and predicates.
package symbols

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/google/mangle/ast"
)

var (
	// Lt is the less-than relation on numbers.
	Lt = ast.PredicateSym{":lt", 2}

	// Le is the less-than-or-equal relation on numbers.
	Le = ast.PredicateSym{":le", 2}

	// MatchPair mode(+, -, -) matches a pair to its elements.
	MatchPair = ast.PredicateSym{":match_pair", 3}

	// MatchCons mode(+, -, -) matches a list to head and tail.
	MatchCons = ast.PredicateSym{":match_cons", 3}

	// MatchNil matches the empty list.
	MatchNil = ast.PredicateSym{":match_nil", 1}

	// MatchEntry mode(+, +, -) matches an entry in a map.
	MatchEntry = ast.PredicateSym{":match_entry", 3}

	// MatchField mode(+, +, -) matches a field in a struct.
	MatchField = ast.PredicateSym{":match_field", 3}

	// ListMember mode(+, -) either checks membership or binds var to every element.
	ListMember = ast.PredicateSym{":list:member", 2}

	// WithinDistance is a relation on numbers X, Y, Z satisfying |X - Y| < Z.
	WithinDistance = ast.PredicateSym{":within_distance", 3}

	// Div is a family of functions mapping X,Y1,.. to (X / Y1) / Y2 ... DIV(X) is 1/x.
	Div = ast.FunctionSym{"fn:div", -1}
	// Mult is a family of functions mapping X,Y1,.. to (X * Y1) * Y2 ... MULT(x) is x.
	Mult = ast.FunctionSym{"fn:mult", -1}
	// Plus is a family of functions mapping X,Y1,.. to (X + Y1) + Y2 ... PLUS(x) is x.
	Plus = ast.FunctionSym{"fn:plus", -1}
	// Minus is a family of functions mapping X,Y1,.. to (X - Y1) - Y2 ...MINUS(x) is -X.
	Minus = ast.FunctionSym{"fn:minus", -1}

	// Collect turns a collection { tuple_1,...tuple_n } into a list [tuple_1, ..., tuple_n].
	Collect = ast.FunctionSym{"fn:collect", -1}
	// CollectDistinct turns a collection { tuple_1,...tuple_n } into a list with distinct elements [tuple_1, ..., tuple_n].
	CollectDistinct = ast.FunctionSym{"fn:collect_distinct", -1}
	// PickAny reduces a set { x_1,...x_n } to a single { x_i },
	PickAny = ast.FunctionSym{"fn:pick_any", 1}
	// Max reduces a set { x_1,...x_n } to { x_i } that is maximal.
	Max = ast.FunctionSym{"fn:max", 1}
	// FloatMax reduces a set of float64 { x_1,...x_n } to { x_i } that is maximal.
	FloatMax = ast.FunctionSym{"fn:float:max", 1}
	// Min reduces a set { x_1,...x_n } to { x_i } that is minimal.
	Min = ast.FunctionSym{"fn:min", 1}
	// FloatMin reduces a set of float64 { x_1,...x_n } to { x_i } that is minimal.
	FloatMin = ast.FunctionSym{"fn:float:min", 1}
	// Sum reduces a set { x_1,...x_n } to { x_1 + ... + x_n }.
	Sum = ast.FunctionSym{"fn:sum", 1}
	// FloatSum reduces a set of float64 { x_1,...x_n } to { x_1 + ... + x_n }.
	FloatSum = ast.FunctionSym{"fn:float:sum", 1}
	// Count reduces a set { x_1,...x_n } to { n }.
	Count = ast.FunctionSym{"fn:count", 0}

	// GroupBy groups all tuples by the values of key variables, e.g. 'group_by(X)'.
	// An empty group_by() treats the whole relation as a group.
	GroupBy = ast.FunctionSym{"fn:group_by", -1}

	// Append appends a element to a list.
	Append = ast.FunctionSym{"fn:list:append", 2}

	// ListGet is a function (List, Number) which returns element at index 'Number'.
	ListGet = ast.FunctionSym{"fn:list:get", 2}

	// ListContains is a function (List, Member) which returns /true if Member is contained in list.
	ListContains = ast.FunctionSym{"fn:list:contains", 2}

	// Len returns length of a list.
	Len = ast.FunctionSym{"fn:list:len", 1}
	// Cons constructs a pair.
	Cons = ast.FunctionSym{"fn:list:cons", 2}
	// Pair constructs a pair.
	Pair = ast.FunctionSym{"fn:pair", 2}
	// MapGet is a function (Map, Key) which returns element at key.
	MapGet = ast.FunctionSym{"fn:map:get", 2}
	// StructGet is a function (Struct, Field) which returns specified field.
	StructGet = ast.FunctionSym{"fn:struct:get", 2}
	// Tuple acts either as identity (one argument), pair (two arguments) or nested pair (more).
	Tuple = ast.FunctionSym{"fn:tuple", -1}
	// List constructs a list.
	List = ast.FunctionSym{"fn:list", -1}
	// Map constructs a map.
	Map = ast.FunctionSym{"fn:map", -1}
	// Struct constructs a struct.
	Struct = ast.FunctionSym{"fn:struct", -1}

	// FunType is a constructor for a function type.
	// fn:Fun(Res, Arg1, ..., ArgN) is Res <= Arg1, ..., ArgN
	FunType = ast.FunctionSym{"fn:Fun", -1}

	// RelType is a constructor for a relation type.
	RelType = ast.FunctionSym{"fn:Rel", -1}

	// PairType is a constructor for a pair type.
	PairType = ast.FunctionSym{"fn:Pair", 2}
	// TupleType is a type-level function that returns a tuple type out of pair types.
	TupleType = ast.FunctionSym{"fn:Tuple", -1}
	// ListType is a constructor for a list type.
	ListType = ast.FunctionSym{"fn:List", 1}
	// MapType is a constructor for a map type.
	MapType = ast.FunctionSym{"fn:Map", 2}
	// StructType is a constructor for a struct type.
	StructType = ast.FunctionSym{"fn:Struct", -1}
	// UnionType is a constructor for a union type.
	UnionType = ast.FunctionSym{"fn:Union", -1}

	// Optional may appear inside StructType to indicate optional fields.
	Optional = ast.FunctionSym{"fn:opt", -1}

	// Package is an improper symbol, used to represent package declaration.
	Package = ast.PredicateSym{"Package", 0}
	// Use is an improper symbol, used to represent use declaration.
	Use = ast.PredicateSym{"Use", 0}

	// TypeConstructors is a list of function symbols used in structured type expressions.
	// Each name is mapped to the corresponding type constructor (a function at the level of types).
	TypeConstructors = map[string]ast.FunctionSym{
		UnionType.Symbol:  UnionType,
		ListType.Symbol:   ListType,
		PairType.Symbol:   PairType,
		TupleType.Symbol:  TupleType,
		MapType.Symbol:    MapType,
		StructType.Symbol: StructType,
		FunType.Symbol:    FunType,
		RelType.Symbol:    RelType,
	}

	// EmptyType is a type without members.
	EmptyType = ast.ApplyFn{UnionType, nil}

	// BuiltinRelations maps each builtin predicate to its argument range list
	BuiltinRelations = map[ast.PredicateSym]ast.BaseTerm{
		// TODO: support float64
		Lt:       NewRelType(ast.NumberBound, ast.NumberBound),
		Le:       NewRelType(ast.NumberBound, ast.NumberBound),
		MatchNil: NewRelType(NewListType(ast.Variable{"X"})),
		MatchCons: NewRelType(
			NewListType(ast.Variable{"X"}), ast.Variable{"X"}, NewListType(ast.Variable{"X"})),
		MatchPair: NewRelType(
			NewPairType(ast.Variable{"X"}, ast.Variable{"Y"}), ast.Variable{"X"}, ast.Variable{"Y"}),
		MatchEntry: NewRelType(
			NewMapType(ast.AnyBound, ast.AnyBound), ast.AnyBound),
		MatchField: NewRelType(
			ast.AnyBound, ast.NameBound, ast.AnyBound),
		ListMember: NewRelType(ast.Variable{"X"}, NewListType(ast.Variable{"X"})),
	}

	errTypeMismatch = errors.New("type mismatch")
)

// NewPairType returns a new PairType.
func NewPairType(left, right ast.BaseTerm) ast.ApplyFn {
	return ast.ApplyFn{PairType, []ast.BaseTerm{left, right}}
}

// NewTupleType returns a new TupleType.
func NewTupleType(parts ...ast.BaseTerm) ast.ApplyFn {
	return ast.ApplyFn{TupleType, parts}
}

// NewListType returns a new ListType.
func NewListType(elem ast.BaseTerm) ast.ApplyFn {
	return ast.ApplyFn{ListType, []ast.BaseTerm{elem}}
}

// NewMapType returns a new MapType.
func NewMapType(left, right ast.BaseTerm) ast.ApplyFn {
	return ast.ApplyFn{MapType, []ast.BaseTerm{left, right}}
}

// NewOpt wraps a label-type pair inside a StructType.
func NewOpt(label, tpe ast.BaseTerm) ast.ApplyFn {
	return ast.ApplyFn{Optional, []ast.BaseTerm{label, tpe}}
}

// NewStructType returns a new StructType.
func NewStructType(args ...ast.BaseTerm) ast.ApplyFn {
	return ast.ApplyFn{StructType, args}
}

// NewFunType returns a new FunType.
func NewFunType(res ast.BaseTerm, args ...ast.BaseTerm) ast.ApplyFn {
	return ast.ApplyFn{FunType, append([]ast.BaseTerm{res}, args...)}
}

// NewRelType returns a new RelType.
func NewRelType(args ...ast.BaseTerm) ast.ApplyFn {
	return ast.ApplyFn{RelType, args}
}

// NewUnionType returns a new UnionType.
func NewUnionType(elems ...ast.BaseTerm) ast.ApplyFn {
	return ast.ApplyFn{UnionType, elems}
}

// IsBaseTypeExpression returns true if c is a base type expression.
func IsBaseTypeExpression(c ast.Constant) bool {
	switch c {
	case ast.AnyBound:
		return true
	case ast.Float64Bound:
		return true
	case ast.NumberBound:
		return true
	case ast.StringBound:
		return true
	default:
		return c.Type == ast.NameType
	}
}

// IsListTypeExpression returns true if tpe is a ListType.
func IsListTypeExpression(tpe ast.BaseTerm) bool {
	listType, ok := tpe.(ast.ApplyFn)
	return ok && listType.Function == ListType
}

// IsMapTypeExpression returns true if tpe is a MapType.
func IsMapTypeExpression(tpe ast.BaseTerm) bool {
	structType, ok := tpe.(ast.ApplyFn)
	return ok && structType.Function.Symbol == MapType.Symbol
}

// IsStructTypeExpression returns true if tpe is a StructType.
func IsStructTypeExpression(tpe ast.BaseTerm) bool {
	structType, ok := tpe.(ast.ApplyFn)
	return ok && structType.Function.Symbol == StructType.Symbol
}

// IsUnionTypeExpression returns true if tpe is a UnionType.
func IsUnionTypeExpression(tpe ast.BaseTerm) bool {
	unionType, ok := tpe.(ast.ApplyFn)
	return ok && unionType.Function.Symbol == UnionType.Symbol
}

// IsRelTypeExpression returns true if tpe is a RelType.
func IsRelTypeExpression(tpe ast.BaseTerm) bool {
	relType, ok := tpe.(ast.ApplyFn)
	return ok && relType.Function == RelType
}

// ListTypeArg returns the type argument of a ListType.
func ListTypeArg(tpe ast.BaseTerm) (ast.BaseTerm, error) {
	listType, ok := tpe.(ast.ApplyFn)
	if !ok || listType.Function != ListType {
		return nil, fmt.Errorf("not a list type expression: %v", tpe)
	}
	if len(listType.Args) != 1 {
		return nil, fmt.Errorf("wrong number of arguments: %v", tpe)
	}
	return listType.Args[0], nil
}

// MapTypeArgs returns the type arguments of a MapType.
func MapTypeArgs(tpe ast.BaseTerm) (ast.BaseTerm, ast.BaseTerm, error) {
	mapType, ok := tpe.(ast.ApplyFn)
	if !ok || mapType.Function != MapType {
		return nil, nil, fmt.Errorf("not a map type expression: %v", tpe)
	}
	if len(mapType.Args) != 2 {
		return nil, nil, fmt.Errorf("wrong number of arguments: %v", tpe)
	}
	return mapType.Args[0], mapType.Args[1], nil
}

// StructTypeArgs returns type arguments of a StructType.
func StructTypeArgs(tpe ast.BaseTerm) ([]ast.BaseTerm, error) {
	structType, ok := tpe.(ast.ApplyFn)
	if !ok || structType.Function.Symbol != StructType.Symbol {
		return nil, fmt.Errorf("not a struct type expression: %v", tpe)
	}
	return structType.Args, nil
}

// IsOptional returns true if an argument of fn:Struct is an optional field.
func IsOptional(structElem ast.BaseTerm) bool {
	opt, ok := structElem.(ast.ApplyFn)
	return ok && opt.Function.Symbol == Optional.Symbol
}

// StructTypeRequiredArgs returns type arguments of a StructType.
func StructTypeRequiredArgs(tpe ast.BaseTerm) ([]ast.BaseTerm, error) {
	structType, ok := tpe.(ast.ApplyFn)
	if !ok || structType.Function.Symbol != StructType.Symbol {
		return nil, fmt.Errorf("not a struct type expression: %v", tpe)
	}
	var required []ast.BaseTerm
	for _, arg := range structType.Args {
		if !IsOptional(arg) {
			required = append(required, arg)
		}
	}
	return required, nil
}

// StructTypeOptionaArgs returns type arguments of a StructType.
func StructTypeOptionaArgs(tpe ast.BaseTerm) ([]ast.BaseTerm, error) {
	structType, ok := tpe.(ast.ApplyFn)
	if !ok || structType.Function.Symbol != StructType.Symbol {
		return nil, fmt.Errorf("not a struct type expression: %v", tpe)
	}
	var optional []ast.BaseTerm
	for _, arg := range structType.Args {
		if IsOptional(arg) {
			optional = append(optional, arg)
		}
	}
	return optional, nil
}

// StructTypeField returns field type for given field.
func StructTypeField(tpe ast.BaseTerm, field ast.Constant) (ast.BaseTerm, error) {
	elems, err := StructTypeArgs(tpe)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(elems); i++ {
		var key ast.BaseTerm
		arg := elems[i]
		if IsOptional(arg) {
			key = arg.(ast.ApplyFn).Args[0]
		} else {
			key = arg
		}
		if key.Equals(field) {
			if IsOptional(arg) {
				return arg.(ast.ApplyFn).Args[1], nil
			}
			i++
			return elems[i], err
		}
	}
	return nil, fmt.Errorf("no field %v in %v", field, tpe)
}

// UnionTypeArgs returns type arguments of a UnionType.
func UnionTypeArgs(tpe ast.BaseTerm) ([]ast.BaseTerm, error) {
	unionType, ok := tpe.(ast.ApplyFn)
	if !ok || unionType.Function.Symbol != UnionType.Symbol {
		return nil, fmt.Errorf("not a union type expression: %v", tpe)
	}
	return unionType.Args, nil
}

// RelTypeArgs returns type arguments of a RelType.
func RelTypeArgs(tpe ast.BaseTerm) ([]ast.BaseTerm, error) {
	relType, ok := tpe.(ast.ApplyFn)
	if !ok || relType.Function != RelType {
		return nil, fmt.Errorf("not a relation type expression: %v", tpe)
	}
	return relType.Args, nil
}

// relTypesFromDecl converts bounds a list of RelTypes.
func relTypesFromDecl(decl ast.Decl) ([]ast.BaseTerm, error) {
	if len(decl.Bounds) == 0 {
		return nil, fmt.Errorf("no bound decls in %v", decl)
	}
	relTypes := make([]ast.BaseTerm, len(decl.Bounds))
	for i, boundDecl := range decl.Bounds {
		relTypes[i] = NewRelType(boundDecl.Bounds...)
	}
	return relTypes, nil
}

// RelTypeFromAlternatives converts list of rel types bounds to union of relation types.
// It is assumed that each alternatives is a RelType.
// An empty list of alternatives corresponds to the empty type fn:Union().
func RelTypeFromAlternatives(alternatives []ast.BaseTerm) ast.BaseTerm {
	if len(alternatives) == 1 {
		return alternatives[0]
	}
	// Could be reduced to a single alternative in some cases.
	return NewUnionType(alternatives...)
}

// RelTypeExprFromDecl converts bounds to relation type expression.
func RelTypeExprFromDecl(decl ast.Decl) (ast.BaseTerm, error) {
	alts, err := relTypesFromDecl(decl)
	if err != nil {
		return nil, err
	}
	return RelTypeFromAlternatives(alts), nil
}

// RelTypeAlternatives converts a relation type expression to a list of alternatives relTypes.
func RelTypeAlternatives(relTypeExpr ast.BaseTerm) []ast.BaseTerm {
	if IsUnionTypeExpression(relTypeExpr) {
		relTypes, _ := UnionTypeArgs(relTypeExpr)
		return relTypes
	}
	return []ast.BaseTerm{relTypeExpr}
}

// TypeHandle provides functionality related to type expression.
type TypeHandle struct {
	expr ast.BaseTerm
}

// NewTypeHandle constructs a TypeHandle.
func NewTypeHandle(expr ast.BaseTerm) (TypeHandle, error) {
	if err := CheckTypeExpression(expr); err != nil {
		return TypeHandle{}, err
	}
	return TypeHandle{expr}, nil
}

// String returns a string represented of this type expression.
func (t TypeHandle) String() string {
	return t.expr.String()
}

// HasType returns true if c has type represented by this TypeHandle.
func (t TypeHandle) HasType(c ast.Constant) bool {
	if baseType, ok := t.expr.(ast.Constant); ok {
		return hasBaseType(baseType, c)
	}
	tpe, ok := t.expr.(ast.ApplyFn)
	if !ok {
		return false // This never happens.
	}
	switch tpe.Function {
	case PairType:
		fst, snd, err := c.PairValue()
		if err != nil {
			return false
		}
		return TypeHandle{tpe.Args[0]}.HasType(fst) &&
			TypeHandle{tpe.Args[1]}.HasType(snd)
	case ListType:
		elementType := TypeHandle{tpe.Args[0]}
		shapeErr, err := c.ListValues(func(e ast.Constant) error {
			if !elementType.HasType(e) {
				return errTypeMismatch
			}
			return nil
		}, func() error {
			return nil
		})
		if shapeErr != nil {
			return false // not a list.
		}
		if errors.Is(err, errTypeMismatch) {
			return false
		}
		return true
	case TupleType:
		return TypeHandle{expandTupleType(tpe.Args)}.HasType(c)
	case MapType:
		if c.IsMapNil() {
			return true
		}
		keyTpe := TypeHandle{tpe.Args[0]}
		valTpe := TypeHandle{tpe.Args[1]}
		e, err := c.MapValues(func(key ast.Constant, val ast.Constant) error {
			if keyTpe.HasType(key) && valTpe.HasType(val) {
				return nil
			}
			return errTypeMismatch
		}, func() error {
			return nil
		})
		return e == nil && err == nil
	case StructType:
		if c.IsStructNil() {
			return len(tpe.Args) == 0
		}
		fieldTpeMap := make(map[ast.Constant]TypeHandle)
		requiredArgs, err := StructTypeRequiredArgs(tpe)
		if err != nil {
			return false
		}
		for i := 0; i < len(requiredArgs); i++ {
			key := requiredArgs[i].(ast.Constant)
			i++
			val := requiredArgs[i]
			fieldTpeMap[key] = TypeHandle{val}
		}
		optArgs, err := StructTypeOptionaArgs(tpe)
		if err != nil {
			return false
		}
		for _, optArg := range optArgs {
			f := optArg.(ast.ApplyFn)
			fieldTpeMap[f.Args[0].(ast.Constant)] = TypeHandle{f.Args[1]}
		}
		seen := make(map[ast.Constant]bool)
		e, err := c.StructValues(func(key ast.Constant, val ast.Constant) error {
			fieldTpe, ok := fieldTpeMap[key]
			if !ok {
				return errTypeMismatch
			}
			seen[key] = true
			if !fieldTpe.HasType(val) {
				return errTypeMismatch
			}
			return nil
		}, func() error {
			return nil
		})
		return e == nil && err == nil && len(fieldTpeMap) == len(seen)
	case UnionType:
		for _, arg := range tpe.Args {
			alt := TypeHandle{arg}
			if alt.HasType(c) {
				return true
			}
		}
		return false
	}
	return false
}

func hasBaseType(typeExpr ast.Constant, c ast.Constant) bool {
	switch typeExpr {
	case ast.AnyBound:
		return true
	case ast.Float64Bound:
		return c.Type == ast.Float64Type
	case ast.NameBound:
		return c.Type == ast.NameType
	case ast.NumberBound:
		return c.Type == ast.NumberType
	case ast.StringBound:
		return c.Type == ast.StringType
	default:
		return typeExpr.Type == ast.NameType && c.Type == ast.NameType && strings.HasPrefix(c.Symbol, typeExpr.Symbol+"/")
	}
}

// CheckTypeExpression returns an error if expr is not a valid type expression.
func CheckTypeExpression(expr ast.BaseTerm) error {
	switch expr := expr.(type) {
	case ast.Constant:
		if !IsBaseTypeExpression(expr) {
			return fmt.Errorf("not a base type expression: %v", expr)
		}
		return nil
	case ast.Variable:
		return fmt.Errorf("not a type expression: %v", expr)
	case ast.ApplyFn:
		fn, ok := TypeConstructors[expr.Function.Symbol]
		if !ok {
			return fmt.Errorf("not a structured type expression: %v", expr)
		}
		args := expr.Args
		if fn == FunType {
			return CheckFunTypeExpression(expr)
		}
		if fn.Arity != -1 && len(args) != fn.Arity {
			return fmt.Errorf("expected %d arguments in type expression %v ", fn.Arity, expr)
		}
		if fn == UnionType && len(args) <= 0 {
			return fmt.Errorf("union type must not be empty %v ", expr)
		}
		if fn == TupleType && len(args) <= 2 {
			return fmt.Errorf("tuple type must have more than 2 args %v ", expr)
		}
		if fn == StructType {
			requiredArgs, err := StructTypeRequiredArgs(expr)
			if err != nil {
				return err
			}
			if len(requiredArgs)%2 != 0 {
				return fmt.Errorf("struct type must have even number of required arguments %v ", expr)
			}
			for i := 0; i < len(requiredArgs); i++ {
				key := requiredArgs[i]
				if c, ok := key.(ast.Constant); !ok || c.Type != ast.NameType {
					return fmt.Errorf("in a struct type expression, odd arguments must be name constants, argument %d (%v) is not %v ", i, key, expr)
				}
				i++
				tpe := requiredArgs[i]
				if err := CheckTypeExpression(tpe); err != nil {
					return fmt.Errorf("in a struct type expression %v : %w", expr, err)
				}
			}
			return nil
		}

		for _, arg := range args {
			if err := CheckTypeExpression(arg); err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("CheckTypeExpression: unexpected case %v %T", expr, expr)
	}
}

// CheckFunTypeExpression checks fun type expression.
func CheckFunTypeExpression(expr ast.ApplyFn) error {
	return fmt.Errorf("CheckFunTypeExpression: not implemented %v", expr)
}

// TypeConforms returns true if left <: right.
func TypeConforms(left ast.BaseTerm, right ast.BaseTerm) bool {
	if left.Equals(right) || right.Equals(ast.AnyBound) {
		return true
	}
	if leftConst, ok := left.(ast.Constant); ok {
		if rightConst, ok := right.(ast.Constant); ok {
			if strings.HasPrefix(leftConst.Symbol, rightConst.Symbol) {
				return true
			}
			return leftConst.Type == ast.NameType && rightConst.Equals(ast.NameBound)
		}
	}
	leftApply, leftApplyOk := left.(ast.ApplyFn)
	rightApply, rightApplyOk := right.(ast.ApplyFn)
	if leftApplyOk && leftApply.Function.Symbol == FunType.Symbol &&
		rightApplyOk && rightApply.Function.Symbol == FunType.Symbol {
		// FunType subtyping is covariant in codomain, contravariant in domain
		// E.g. /genus_species <: /name and /animal/bird <: /animal
		// therefore FunType(/genus_species <= /animal) <: FunType(/name <= /animal/bird)
		leftCodomain, rightCodomain := leftApply.Args[0], rightApply.Args[0]
		if !TypeConforms(leftCodomain, rightCodomain) {
			return false
		}
		leftDomain, rightDomain := leftApply.Args[1:], rightApply.Args[1:]

		for i, leftArg := range leftDomain {
			if !TypeConforms(rightDomain[i], leftArg) {
				return false
			}
		}
		return true
	}

	if leftApplyOk && leftApply.Function.Symbol == UnionType.Symbol {
		for _, leftItem := range leftApply.Args {
			if !TypeConforms(leftItem, right) {
				return false
			}
		}
		return true
	}
	if rightApplyOk && rightApply.Function.Symbol == UnionType.Symbol {
		for _, rightItem := range rightApply.Args {
			if TypeConforms(left, rightItem) {
				return true
			}
		}
	}
	if leftApplyOk && leftApply.Function.Symbol == ListType.Symbol {
		if rightApplyOk && rightApply.Function.Symbol == ListType.Symbol {
			return TypeConforms(leftApply.Args[0], rightApply.Args[0])
		}
	}
	if leftApplyOk && leftApply.Function.Symbol == MapType.Symbol {
		if rightApplyOk && rightApply.Function.Symbol == MapType.Symbol {
			return TypeConforms(rightApply.Args[0], leftApply.Args[0]) && TypeConforms(leftApply.Args[1], rightApply.Args[1])
		}
	}
	if leftApplyOk && leftApply.Function.Symbol == StructType.Symbol {
		if rightApplyOk && rightApply.Function.Symbol == StructType.Symbol {
			leftRequired, err := StructTypeRequiredArgs(left)
			if err != nil {
				return false
			}
			rightRequired, err := StructTypeRequiredArgs(right)
			if err != nil {
				return false
			}
			if len(leftRequired) < len(rightRequired) {
				return false
			}
			leftMap := make(map[string]ast.BaseTerm)
			for i := 0; i < len(leftRequired); i++ {
				leftKey, _ := leftRequired[i].(ast.Constant)
				i++
				leftMap[leftKey.Symbol] = leftRequired[i]
			}

			for j := 0; j < len(rightRequired); j++ {
				rightKey, _ := rightRequired[j].(ast.Constant)
				j++
				rightTpe := rightRequired[j]
				leftTpe, ok := leftMap[rightKey.Symbol]
				if !ok || !TypeConforms(leftTpe, rightTpe) {
					return false
				}
			}
			leftOpt, err := StructTypeOptionaArgs(left)
			if err != nil {
				return false
			}
			rightOpt, err := StructTypeOptionaArgs(right)
			if err != nil {
				return false
			}
			leftOptMap := make(map[string]ast.BaseTerm)
			for _, opt := range leftOpt {
				optApply, ok := opt.(ast.ApplyFn)
				if !ok {
					return false
				}
				leftOptMap[optApply.Args[0].(ast.Constant).Symbol] = optApply.Args[1]
			}
			for _, opt := range rightOpt {
				optApply, ok := opt.(ast.ApplyFn)
				if !ok {
					return false
				}
				key := optApply.Args[0].(ast.Constant).Symbol
				rightTpe := optApply.Args[1]
				leftTpe, ok := leftMap[key]
				if ok && !TypeConforms(leftTpe, rightTpe) {
					return false
				}
				if !ok {
					leftTpe, ok := leftOptMap[key]
					if ok && !TypeConforms(leftTpe, rightTpe) {
						return false
					}
				}
			}
			return true
		}
	}

	if leftApplyOk && leftApply.Function.Symbol == PairType.Symbol {
		if rightApplyOk && rightApply.Function.Symbol == PairType.Symbol {
			return TypeConforms(leftApply.Args[0], rightApply.Args[0]) && TypeConforms(leftApply.Args[1], rightApply.Args[1])
		}
	}
	if leftTuple, ok := left.(ast.ApplyFn); ok && leftTuple.Function.Symbol == TupleType.Symbol {
		if rightTuple, ok := right.(ast.ApplyFn); ok && rightTuple.Function.Symbol == TupleType.Symbol {
			for i, leftArg := range leftTuple.Args {
				if !TypeConforms(leftArg, rightTuple.Args[i]) {
					return false
				}
			}
			return true
		}
	}
	if leftTuple, ok := left.(ast.ApplyFn); ok && leftTuple.Function.Symbol == RelType.Symbol {
		if rightTuple, ok := right.(ast.ApplyFn); ok && rightTuple.Function.Symbol == RelType.Symbol {
			for i, leftArg := range leftTuple.Args {
				if !TypeConforms(leftArg, rightTuple.Args[i]) {
					return false
				}
			}
			return true
		}
	}

	return false
}

func expandTupleType(args []ast.BaseTerm) ast.BaseTerm {
	res := NewPairType(args[len(args)-2], args[len(args)-1])
	for j := len(args) - 3; j >= 0; j-- {
		res = NewPairType(args[j], res)
	}
	return res
}

// UpperBound returns upper bound of type expressions.
func UpperBound(typeExprs []ast.BaseTerm) ast.BaseTerm {
	var worklist []ast.BaseTerm
	for _, typeExpr := range typeExprs {
		if ast.AnyBound.Equals(typeExpr) {
			return ast.AnyBound
		}
		if union, ok := typeExpr.(ast.ApplyFn); ok && union.Function == UnionType {
			worklist = append(worklist, union.Args...)
			continue
		}
		worklist = append(worklist, typeExpr)
	}
	if len(worklist) == 0 {
		return EmptyType
	}
	reduced := []ast.BaseTerm{worklist[0]}
	worklist = worklist[1:]
typeExprLoop:
	for _, typeExpr := range worklist {
		for i, existing := range reduced {
			if TypeConforms(typeExpr, existing) {
				continue typeExprLoop
			}
			if TypeConforms(existing, typeExpr) {
				reduced[i] = typeExpr
				continue typeExprLoop
			}
		}
		reduced = append(reduced, typeExpr)
	}
	if len(reduced) == 1 {
		return reduced[0]
	}
	sort.Slice(reduced, func(i, j int) bool { return reduced[i].Hash() < reduced[j].Hash() })
	return ast.ApplyFn{UnionType, reduced}
}

func intersectType(a, b ast.BaseTerm) ast.BaseTerm {
	if a.Equals(b) {
		return a
	}
	if a.Equals(ast.AnyBound) {
		return b
	}
	if b.Equals(ast.AnyBound) {
		return a
	}
	if TypeConforms(a, b) {
		return a
	}
	if TypeConforms(b, a) {
		return b
	}
	if aUnion, ok := a.(ast.ApplyFn); ok && aUnion.Function == UnionType {
		var res []ast.BaseTerm
		for _, elem := range aUnion.Args {
			if u := intersectType(elem, b); !u.Equals(EmptyType) {
				res = append(res, u)
			}
		}
		return UpperBound(res)
	}
	if bUnion, ok := b.(ast.ApplyFn); ok && bUnion.Function == UnionType {
		var res []ast.BaseTerm
		for _, elem := range bUnion.Args {
			if TypeConforms(a, elem) {
				res = append(res, a)
			} else if TypeConforms(elem, a) {
				res = append(res, elem)
			}
		}
		return UpperBound(res)
	}

	return EmptyType
}

// LowerBound returns a lower bound of type expressions.
func LowerBound(typeExprs []ast.BaseTerm) ast.BaseTerm {
	var typeExpr ast.BaseTerm = ast.AnyBound
	for _, t := range typeExprs {
		if typeExpr = intersectType(typeExpr, t); typeExpr.Equals(EmptyType) {
			return EmptyType
		}
	}
	return typeExpr
}
