// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package gen

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MangleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var manglelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func manglelexerLexerInit() {
	staticData := &manglelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'.'", "'descr'", "'inclusion'", "':'", "'{'", "'}'", "", "", "'Package'",
		"'Use'", "'Decl'", "'bound'", "'let'", "'do'", "'('", "')'", "'['",
		"']'", "'='", "'!='", "','", "'!'", "'<'", "'<='", "'>'", "'>='", "':-'",
		"'\\n'", "'|>'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "WHITESPACE", "COMMENT", "PACKAGE", "USE",
		"DECL", "BOUND", "LET", "DO", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET",
		"EQ", "BANGEQ", "COMMA", "BANG", "LESS", "LESSEQ", "GREATER", "GREATEREQ",
		"COLONDASH", "NEWLINE", "PIPEGREATER", "NUMBER", "FLOAT", "VARIABLE",
		"NAME", "CONSTANT", "STRING",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "WHITESPACE", "COMMENT",
		"PACKAGE", "USE", "DECL", "BOUND", "LET", "DO", "LPAREN", "RPAREN",
		"LBRACKET", "RBRACKET", "EQ", "BANGEQ", "COMMA", "BANG", "LESS", "LESSEQ",
		"GREATER", "GREATEREQ", "COLONDASH", "NEWLINE", "PIPEGREATER", "LETTER",
		"DIGIT", "NUMBER", "FLOAT", "EXPONENT", "VARIABLE_START", "VARIABLE_CHAR",
		"VARIABLE", "NAME_CHAR", "NAME", "CONSTANT_CHAR", "CONSTANT", "STRING",
		"SHORT_STRING", "LONG_STRING", "LONG_STRING_ITEM", "LONG_STRING_CHAR",
		"STRING_ESCAPE_SEQ",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 354, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 4, 6, 121, 8, 6, 11, 6, 12, 6, 122, 1, 6, 1, 6, 1, 7, 1, 7, 5, 7,
		129, 8, 7, 10, 7, 12, 7, 132, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 31, 3, 31, 206, 8, 31, 1, 31, 1, 31, 5, 31, 210, 8, 31, 10, 31,
		12, 31, 213, 9, 31, 1, 32, 3, 32, 216, 8, 32, 1, 32, 4, 32, 219, 8, 32,
		11, 32, 12, 32, 220, 1, 32, 1, 32, 4, 32, 225, 8, 32, 11, 32, 12, 32, 226,
		1, 32, 3, 32, 230, 8, 32, 1, 32, 3, 32, 233, 8, 32, 1, 32, 1, 32, 4, 32,
		237, 8, 32, 11, 32, 12, 32, 238, 1, 32, 3, 32, 242, 8, 32, 3, 32, 244,
		8, 32, 1, 33, 1, 33, 3, 33, 248, 8, 33, 1, 33, 4, 33, 251, 8, 33, 11, 33,
		12, 33, 252, 1, 34, 1, 34, 1, 35, 1, 35, 3, 35, 259, 8, 35, 1, 36, 1, 36,
		1, 36, 5, 36, 264, 8, 36, 10, 36, 12, 36, 267, 9, 36, 3, 36, 269, 8, 36,
		1, 37, 1, 37, 1, 37, 3, 37, 274, 8, 37, 1, 38, 3, 38, 277, 8, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 5, 38, 283, 8, 38, 10, 38, 12, 38, 286, 9, 38, 1,
		39, 1, 39, 1, 39, 3, 39, 291, 8, 39, 1, 40, 1, 40, 4, 40, 295, 8, 40, 11,
		40, 12, 40, 296, 1, 40, 1, 40, 4, 40, 301, 8, 40, 11, 40, 12, 40, 302,
		5, 40, 305, 8, 40, 10, 40, 12, 40, 308, 9, 40, 1, 41, 1, 41, 3, 41, 312,
		8, 41, 1, 42, 1, 42, 1, 42, 5, 42, 317, 8, 42, 10, 42, 12, 42, 320, 9,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 5, 42, 326, 8, 42, 10, 42, 12, 42, 329,
		9, 42, 1, 42, 3, 42, 332, 8, 42, 1, 43, 1, 43, 5, 43, 336, 8, 43, 10, 43,
		12, 43, 339, 9, 43, 1, 43, 1, 43, 1, 44, 1, 44, 3, 44, 345, 8, 44, 1, 45,
		1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 353, 8, 46, 1, 337, 0, 47, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 0, 61, 0, 63, 30, 65, 31, 67, 0, 69, 0, 71, 0, 73, 32, 75, 0, 77, 33,
		79, 0, 81, 34, 83, 35, 85, 0, 87, 0, 89, 0, 91, 0, 93, 0, 1, 0, 10, 3,
		0, 9, 10, 12, 13, 32, 32, 1, 0, 10, 10, 2, 0, 65, 90, 97, 122, 2, 0, 69,
		69, 101, 101, 2, 0, 43, 43, 45, 45, 2, 0, 58, 58, 95, 95, 4, 0, 37, 37,
		45, 46, 95, 95, 126, 126, 4, 0, 10, 10, 12, 13, 39, 39, 92, 92, 4, 0, 10,
		10, 12, 13, 34, 34, 92, 92, 1, 0, 92, 92, 377, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0,
		77, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 1, 95, 1, 0, 0, 0,
		3, 97, 1, 0, 0, 0, 5, 103, 1, 0, 0, 0, 7, 113, 1, 0, 0, 0, 9, 115, 1, 0,
		0, 0, 11, 117, 1, 0, 0, 0, 13, 120, 1, 0, 0, 0, 15, 126, 1, 0, 0, 0, 17,
		135, 1, 0, 0, 0, 19, 143, 1, 0, 0, 0, 21, 147, 1, 0, 0, 0, 23, 152, 1,
		0, 0, 0, 25, 158, 1, 0, 0, 0, 27, 162, 1, 0, 0, 0, 29, 165, 1, 0, 0, 0,
		31, 167, 1, 0, 0, 0, 33, 169, 1, 0, 0, 0, 35, 171, 1, 0, 0, 0, 37, 173,
		1, 0, 0, 0, 39, 175, 1, 0, 0, 0, 41, 178, 1, 0, 0, 0, 43, 180, 1, 0, 0,
		0, 45, 182, 1, 0, 0, 0, 47, 184, 1, 0, 0, 0, 49, 187, 1, 0, 0, 0, 51, 189,
		1, 0, 0, 0, 53, 192, 1, 0, 0, 0, 55, 195, 1, 0, 0, 0, 57, 197, 1, 0, 0,
		0, 59, 200, 1, 0, 0, 0, 61, 202, 1, 0, 0, 0, 63, 205, 1, 0, 0, 0, 65, 243,
		1, 0, 0, 0, 67, 245, 1, 0, 0, 0, 69, 254, 1, 0, 0, 0, 71, 258, 1, 0, 0,
		0, 73, 268, 1, 0, 0, 0, 75, 273, 1, 0, 0, 0, 77, 276, 1, 0, 0, 0, 79, 290,
		1, 0, 0, 0, 81, 292, 1, 0, 0, 0, 83, 311, 1, 0, 0, 0, 85, 331, 1, 0, 0,
		0, 87, 333, 1, 0, 0, 0, 89, 344, 1, 0, 0, 0, 91, 346, 1, 0, 0, 0, 93, 352,
		1, 0, 0, 0, 95, 96, 5, 46, 0, 0, 96, 2, 1, 0, 0, 0, 97, 98, 5, 100, 0,
		0, 98, 99, 5, 101, 0, 0, 99, 100, 5, 115, 0, 0, 100, 101, 5, 99, 0, 0,
		101, 102, 5, 114, 0, 0, 102, 4, 1, 0, 0, 0, 103, 104, 5, 105, 0, 0, 104,
		105, 5, 110, 0, 0, 105, 106, 5, 99, 0, 0, 106, 107, 5, 108, 0, 0, 107,
		108, 5, 117, 0, 0, 108, 109, 5, 115, 0, 0, 109, 110, 5, 105, 0, 0, 110,
		111, 5, 111, 0, 0, 111, 112, 5, 110, 0, 0, 112, 6, 1, 0, 0, 0, 113, 114,
		5, 58, 0, 0, 114, 8, 1, 0, 0, 0, 115, 116, 5, 123, 0, 0, 116, 10, 1, 0,
		0, 0, 117, 118, 5, 125, 0, 0, 118, 12, 1, 0, 0, 0, 119, 121, 7, 0, 0, 0,
		120, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 6, 6, 0, 0, 125, 14, 1,
		0, 0, 0, 126, 130, 5, 35, 0, 0, 127, 129, 8, 1, 0, 0, 128, 127, 1, 0, 0,
		0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		133, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133, 134, 6, 7, 0, 0, 134, 16, 1,
		0, 0, 0, 135, 136, 5, 80, 0, 0, 136, 137, 5, 97, 0, 0, 137, 138, 5, 99,
		0, 0, 138, 139, 5, 107, 0, 0, 139, 140, 5, 97, 0, 0, 140, 141, 5, 103,
		0, 0, 141, 142, 5, 101, 0, 0, 142, 18, 1, 0, 0, 0, 143, 144, 5, 85, 0,
		0, 144, 145, 5, 115, 0, 0, 145, 146, 5, 101, 0, 0, 146, 20, 1, 0, 0, 0,
		147, 148, 5, 68, 0, 0, 148, 149, 5, 101, 0, 0, 149, 150, 5, 99, 0, 0, 150,
		151, 5, 108, 0, 0, 151, 22, 1, 0, 0, 0, 152, 153, 5, 98, 0, 0, 153, 154,
		5, 111, 0, 0, 154, 155, 5, 117, 0, 0, 155, 156, 5, 110, 0, 0, 156, 157,
		5, 100, 0, 0, 157, 24, 1, 0, 0, 0, 158, 159, 5, 108, 0, 0, 159, 160, 5,
		101, 0, 0, 160, 161, 5, 116, 0, 0, 161, 26, 1, 0, 0, 0, 162, 163, 5, 100,
		0, 0, 163, 164, 5, 111, 0, 0, 164, 28, 1, 0, 0, 0, 165, 166, 5, 40, 0,
		0, 166, 30, 1, 0, 0, 0, 167, 168, 5, 41, 0, 0, 168, 32, 1, 0, 0, 0, 169,
		170, 5, 91, 0, 0, 170, 34, 1, 0, 0, 0, 171, 172, 5, 93, 0, 0, 172, 36,
		1, 0, 0, 0, 173, 174, 5, 61, 0, 0, 174, 38, 1, 0, 0, 0, 175, 176, 5, 33,
		0, 0, 176, 177, 5, 61, 0, 0, 177, 40, 1, 0, 0, 0, 178, 179, 5, 44, 0, 0,
		179, 42, 1, 0, 0, 0, 180, 181, 5, 33, 0, 0, 181, 44, 1, 0, 0, 0, 182, 183,
		5, 60, 0, 0, 183, 46, 1, 0, 0, 0, 184, 185, 5, 60, 0, 0, 185, 186, 5, 61,
		0, 0, 186, 48, 1, 0, 0, 0, 187, 188, 5, 62, 0, 0, 188, 50, 1, 0, 0, 0,
		189, 190, 5, 62, 0, 0, 190, 191, 5, 61, 0, 0, 191, 52, 1, 0, 0, 0, 192,
		193, 5, 58, 0, 0, 193, 194, 5, 45, 0, 0, 194, 54, 1, 0, 0, 0, 195, 196,
		5, 10, 0, 0, 196, 56, 1, 0, 0, 0, 197, 198, 5, 124, 0, 0, 198, 199, 5,
		62, 0, 0, 199, 58, 1, 0, 0, 0, 200, 201, 7, 2, 0, 0, 201, 60, 1, 0, 0,
		0, 202, 203, 2, 48, 57, 0, 203, 62, 1, 0, 0, 0, 204, 206, 5, 45, 0, 0,
		205, 204, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207,
		211, 3, 61, 30, 0, 208, 210, 3, 61, 30, 0, 209, 208, 1, 0, 0, 0, 210, 213,
		1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 64, 1, 0,
		0, 0, 213, 211, 1, 0, 0, 0, 214, 216, 5, 45, 0, 0, 215, 214, 1, 0, 0, 0,
		215, 216, 1, 0, 0, 0, 216, 218, 1, 0, 0, 0, 217, 219, 3, 61, 30, 0, 218,
		217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 220, 221,
		1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 224, 5, 46, 0, 0, 223, 225, 3, 61,
		30, 0, 224, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0,
		226, 227, 1, 0, 0, 0, 227, 229, 1, 0, 0, 0, 228, 230, 3, 67, 33, 0, 229,
		228, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 244, 1, 0, 0, 0, 231, 233,
		5, 45, 0, 0, 232, 231, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234, 1, 0,
		0, 0, 234, 236, 5, 46, 0, 0, 235, 237, 3, 61, 30, 0, 236, 235, 1, 0, 0,
		0, 237, 238, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239,
		241, 1, 0, 0, 0, 240, 242, 3, 67, 33, 0, 241, 240, 1, 0, 0, 0, 241, 242,
		1, 0, 0, 0, 242, 244, 1, 0, 0, 0, 243, 215, 1, 0, 0, 0, 243, 232, 1, 0,
		0, 0, 244, 66, 1, 0, 0, 0, 245, 247, 7, 3, 0, 0, 246, 248, 7, 4, 0, 0,
		247, 246, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 1, 0, 0, 0, 249,
		251, 3, 61, 30, 0, 250, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 250,
		1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 68, 1, 0, 0, 0, 254, 255, 2, 65,
		90, 0, 255, 70, 1, 0, 0, 0, 256, 259, 3, 59, 29, 0, 257, 259, 3, 61, 30,
		0, 258, 256, 1, 0, 0, 0, 258, 257, 1, 0, 0, 0, 259, 72, 1, 0, 0, 0, 260,
		269, 5, 95, 0, 0, 261, 265, 3, 69, 34, 0, 262, 264, 3, 71, 35, 0, 263,
		262, 1, 0, 0, 0, 264, 267, 1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 265, 266,
		1, 0, 0, 0, 266, 269, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 268, 260, 1, 0,
		0, 0, 268, 261, 1, 0, 0, 0, 269, 74, 1, 0, 0, 0, 270, 274, 3, 59, 29, 0,
		271, 274, 3, 61, 30, 0, 272, 274, 7, 5, 0, 0, 273, 270, 1, 0, 0, 0, 273,
		271, 1, 0, 0, 0, 273, 272, 1, 0, 0, 0, 274, 76, 1, 0, 0, 0, 275, 277, 5,
		58, 0, 0, 276, 275, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 1, 0, 0,
		0, 278, 284, 2, 97, 122, 0, 279, 283, 3, 75, 37, 0, 280, 281, 5, 46, 0,
		0, 281, 283, 3, 75, 37, 0, 282, 279, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0,
		283, 286, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285,
		78, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 287, 291, 3, 59, 29, 0, 288, 291,
		3, 61, 30, 0, 289, 291, 7, 6, 0, 0, 290, 287, 1, 0, 0, 0, 290, 288, 1,
		0, 0, 0, 290, 289, 1, 0, 0, 0, 291, 80, 1, 0, 0, 0, 292, 294, 5, 47, 0,
		0, 293, 295, 3, 79, 39, 0, 294, 293, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0,
		296, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 306, 1, 0, 0, 0, 298,
		300, 5, 47, 0, 0, 299, 301, 3, 79, 39, 0, 300, 299, 1, 0, 0, 0, 301, 302,
		1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 305, 1, 0,
		0, 0, 304, 298, 1, 0, 0, 0, 305, 308, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0,
		306, 307, 1, 0, 0, 0, 307, 82, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 309, 312,
		3, 85, 42, 0, 310, 312, 3, 87, 43, 0, 311, 309, 1, 0, 0, 0, 311, 310, 1,
		0, 0, 0, 312, 84, 1, 0, 0, 0, 313, 318, 5, 39, 0, 0, 314, 317, 3, 93, 46,
		0, 315, 317, 8, 7, 0, 0, 316, 314, 1, 0, 0, 0, 316, 315, 1, 0, 0, 0, 317,
		320, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 321,
		1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 321, 332, 5, 39, 0, 0, 322, 327, 5, 34,
		0, 0, 323, 326, 3, 93, 46, 0, 324, 326, 8, 8, 0, 0, 325, 323, 1, 0, 0,
		0, 325, 324, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 327,
		328, 1, 0, 0, 0, 328, 330, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 330, 332,
		5, 34, 0, 0, 331, 313, 1, 0, 0, 0, 331, 322, 1, 0, 0, 0, 332, 86, 1, 0,
		0, 0, 333, 337, 5, 96, 0, 0, 334, 336, 3, 89, 44, 0, 335, 334, 1, 0, 0,
		0, 336, 339, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 338,
		340, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 340, 341, 5, 96, 0, 0, 341, 88,
		1, 0, 0, 0, 342, 345, 3, 91, 45, 0, 343, 345, 3, 93, 46, 0, 344, 342, 1,
		0, 0, 0, 344, 343, 1, 0, 0, 0, 345, 90, 1, 0, 0, 0, 346, 347, 8, 9, 0,
		0, 347, 92, 1, 0, 0, 0, 348, 349, 5, 92, 0, 0, 349, 353, 9, 0, 0, 0, 350,
		351, 5, 92, 0, 0, 351, 353, 3, 55, 27, 0, 352, 348, 1, 0, 0, 0, 352, 350,
		1, 0, 0, 0, 353, 94, 1, 0, 0, 0, 35, 0, 122, 130, 205, 211, 215, 220, 226,
		229, 232, 238, 241, 243, 247, 252, 258, 265, 268, 273, 276, 282, 284, 290,
		296, 302, 306, 311, 316, 318, 325, 327, 331, 337, 344, 352, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MangleLexerInit initializes any static state used to implement MangleLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMangleLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MangleLexerInit() {
	staticData := &manglelexerLexerStaticData
	staticData.once.Do(manglelexerLexerInit)
}

// NewMangleLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMangleLexer(input antlr.CharStream) *MangleLexer {
	MangleLexerInit()
	l := new(MangleLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &manglelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Mangle.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MangleLexer tokens.
const (
	MangleLexerT__0        = 1
	MangleLexerT__1        = 2
	MangleLexerT__2        = 3
	MangleLexerT__3        = 4
	MangleLexerT__4        = 5
	MangleLexerT__5        = 6
	MangleLexerWHITESPACE  = 7
	MangleLexerCOMMENT     = 8
	MangleLexerPACKAGE     = 9
	MangleLexerUSE         = 10
	MangleLexerDECL        = 11
	MangleLexerBOUND       = 12
	MangleLexerLET         = 13
	MangleLexerDO          = 14
	MangleLexerLPAREN      = 15
	MangleLexerRPAREN      = 16
	MangleLexerLBRACKET    = 17
	MangleLexerRBRACKET    = 18
	MangleLexerEQ          = 19
	MangleLexerBANGEQ      = 20
	MangleLexerCOMMA       = 21
	MangleLexerBANG        = 22
	MangleLexerLESS        = 23
	MangleLexerLESSEQ      = 24
	MangleLexerGREATER     = 25
	MangleLexerGREATEREQ   = 26
	MangleLexerCOLONDASH   = 27
	MangleLexerNEWLINE     = 28
	MangleLexerPIPEGREATER = 29
	MangleLexerNUMBER      = 30
	MangleLexerFLOAT       = 31
	MangleLexerVARIABLE    = 32
	MangleLexerNAME        = 33
	MangleLexerCONSTANT    = 34
	MangleLexerSTRING      = 35
)