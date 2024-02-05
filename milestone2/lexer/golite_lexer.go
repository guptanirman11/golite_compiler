// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package lexer

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

type GoliteLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var golitelexerLexerStaticData struct {
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

func golitelexerLexerInit() {
	staticData := &golitelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'printf'", "'struct'", "'true'", "'false'", "'nil'", "'new'", "'int'",
		"'bool'", "'scan'", "'delete'", "'if'", "'else'", "'for'", "'return'",
		"'type'", "'var'", "'func'", "", "", "'('", "')'", "'{'", "'}'", "'='",
		"'>'", "'>='", "'<'", "'<='", "'!'", "'=='", "'!='", "'&&'", "'||'",
		"'+'", "'-'", "'/'", "'*'", "':'", "", "';'", "'.'", "','",
	}
	staticData.symbolicNames = []string{
		"", "PRINT", "STRUCT", "TRUE", "FALSE", "NIL", "NEW", "INT", "BOOL",
		"SCAN", "DELETE", "IF", "ELSE", "FOR", "RETURN", "TYPE", "VAR", "FUNC",
		"NUMBER", "STRING", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "EQUALS",
		"GTN", "GETN", "LTN", "LETN", "NOT", "EQ", "NEQ", "AND", "OR", "PLUS",
		"MINUS", "FSLASH", "ASTERISK", "COLON", "IDENTIFIER", "SEMICOLON", "DOT",
		"COMMA", "WS", "COMMENT",
	}
	staticData.ruleNames = []string{
		"PRINT", "STRUCT", "TRUE", "FALSE", "NIL", "NEW", "INT", "BOOL", "SCAN",
		"DELETE", "IF", "ELSE", "FOR", "RETURN", "TYPE", "VAR", "FUNC", "NUMBER",
		"STRING", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "EQUALS", "GTN", "GETN",
		"LTN", "LETN", "NOT", "EQ", "NEQ", "AND", "OR", "PLUS", "MINUS", "FSLASH",
		"ASTERISK", "COLON", "IDENTIFIER", "SEMICOLON", "DOT", "COMMA", "WS",
		"COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 268, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 17, 3, 17, 178, 8, 17, 1, 17, 4, 17, 181, 8, 17, 11, 17, 12,
		17, 182, 1, 18, 1, 18, 5, 18, 187, 8, 18, 10, 18, 12, 18, 190, 9, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 5, 38, 240, 8, 38, 10, 38, 12, 38,
		243, 9, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 4, 42, 252,
		8, 42, 11, 42, 12, 42, 253, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 5,
		43, 262, 8, 43, 10, 43, 12, 43, 265, 9, 43, 1, 43, 1, 43, 0, 0, 44, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38,
		77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 1, 0, 7, 2, 0, 43, 43,
		45, 45, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0,
		48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10,
		10, 13, 13, 273, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0,
		0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0,
		0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0,
		0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1,
		0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83,
		1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 1, 89, 1, 0, 0, 0, 3,
		96, 1, 0, 0, 0, 5, 103, 1, 0, 0, 0, 7, 108, 1, 0, 0, 0, 9, 114, 1, 0, 0,
		0, 11, 118, 1, 0, 0, 0, 13, 122, 1, 0, 0, 0, 15, 126, 1, 0, 0, 0, 17, 131,
		1, 0, 0, 0, 19, 136, 1, 0, 0, 0, 21, 143, 1, 0, 0, 0, 23, 146, 1, 0, 0,
		0, 25, 151, 1, 0, 0, 0, 27, 155, 1, 0, 0, 0, 29, 162, 1, 0, 0, 0, 31, 167,
		1, 0, 0, 0, 33, 171, 1, 0, 0, 0, 35, 177, 1, 0, 0, 0, 37, 184, 1, 0, 0,
		0, 39, 193, 1, 0, 0, 0, 41, 195, 1, 0, 0, 0, 43, 197, 1, 0, 0, 0, 45, 199,
		1, 0, 0, 0, 47, 201, 1, 0, 0, 0, 49, 203, 1, 0, 0, 0, 51, 205, 1, 0, 0,
		0, 53, 208, 1, 0, 0, 0, 55, 210, 1, 0, 0, 0, 57, 213, 1, 0, 0, 0, 59, 215,
		1, 0, 0, 0, 61, 218, 1, 0, 0, 0, 63, 221, 1, 0, 0, 0, 65, 224, 1, 0, 0,
		0, 67, 227, 1, 0, 0, 0, 69, 229, 1, 0, 0, 0, 71, 231, 1, 0, 0, 0, 73, 233,
		1, 0, 0, 0, 75, 235, 1, 0, 0, 0, 77, 237, 1, 0, 0, 0, 79, 244, 1, 0, 0,
		0, 81, 246, 1, 0, 0, 0, 83, 248, 1, 0, 0, 0, 85, 251, 1, 0, 0, 0, 87, 257,
		1, 0, 0, 0, 89, 90, 5, 112, 0, 0, 90, 91, 5, 114, 0, 0, 91, 92, 5, 105,
		0, 0, 92, 93, 5, 110, 0, 0, 93, 94, 5, 116, 0, 0, 94, 95, 5, 102, 0, 0,
		95, 2, 1, 0, 0, 0, 96, 97, 5, 115, 0, 0, 97, 98, 5, 116, 0, 0, 98, 99,
		5, 114, 0, 0, 99, 100, 5, 117, 0, 0, 100, 101, 5, 99, 0, 0, 101, 102, 5,
		116, 0, 0, 102, 4, 1, 0, 0, 0, 103, 104, 5, 116, 0, 0, 104, 105, 5, 114,
		0, 0, 105, 106, 5, 117, 0, 0, 106, 107, 5, 101, 0, 0, 107, 6, 1, 0, 0,
		0, 108, 109, 5, 102, 0, 0, 109, 110, 5, 97, 0, 0, 110, 111, 5, 108, 0,
		0, 111, 112, 5, 115, 0, 0, 112, 113, 5, 101, 0, 0, 113, 8, 1, 0, 0, 0,
		114, 115, 5, 110, 0, 0, 115, 116, 5, 105, 0, 0, 116, 117, 5, 108, 0, 0,
		117, 10, 1, 0, 0, 0, 118, 119, 5, 110, 0, 0, 119, 120, 5, 101, 0, 0, 120,
		121, 5, 119, 0, 0, 121, 12, 1, 0, 0, 0, 122, 123, 5, 105, 0, 0, 123, 124,
		5, 110, 0, 0, 124, 125, 5, 116, 0, 0, 125, 14, 1, 0, 0, 0, 126, 127, 5,
		98, 0, 0, 127, 128, 5, 111, 0, 0, 128, 129, 5, 111, 0, 0, 129, 130, 5,
		108, 0, 0, 130, 16, 1, 0, 0, 0, 131, 132, 5, 115, 0, 0, 132, 133, 5, 99,
		0, 0, 133, 134, 5, 97, 0, 0, 134, 135, 5, 110, 0, 0, 135, 18, 1, 0, 0,
		0, 136, 137, 5, 100, 0, 0, 137, 138, 5, 101, 0, 0, 138, 139, 5, 108, 0,
		0, 139, 140, 5, 101, 0, 0, 140, 141, 5, 116, 0, 0, 141, 142, 5, 101, 0,
		0, 142, 20, 1, 0, 0, 0, 143, 144, 5, 105, 0, 0, 144, 145, 5, 102, 0, 0,
		145, 22, 1, 0, 0, 0, 146, 147, 5, 101, 0, 0, 147, 148, 5, 108, 0, 0, 148,
		149, 5, 115, 0, 0, 149, 150, 5, 101, 0, 0, 150, 24, 1, 0, 0, 0, 151, 152,
		5, 102, 0, 0, 152, 153, 5, 111, 0, 0, 153, 154, 5, 114, 0, 0, 154, 26,
		1, 0, 0, 0, 155, 156, 5, 114, 0, 0, 156, 157, 5, 101, 0, 0, 157, 158, 5,
		116, 0, 0, 158, 159, 5, 117, 0, 0, 159, 160, 5, 114, 0, 0, 160, 161, 5,
		110, 0, 0, 161, 28, 1, 0, 0, 0, 162, 163, 5, 116, 0, 0, 163, 164, 5, 121,
		0, 0, 164, 165, 5, 112, 0, 0, 165, 166, 5, 101, 0, 0, 166, 30, 1, 0, 0,
		0, 167, 168, 5, 118, 0, 0, 168, 169, 5, 97, 0, 0, 169, 170, 5, 114, 0,
		0, 170, 32, 1, 0, 0, 0, 171, 172, 5, 102, 0, 0, 172, 173, 5, 117, 0, 0,
		173, 174, 5, 110, 0, 0, 174, 175, 5, 99, 0, 0, 175, 34, 1, 0, 0, 0, 176,
		178, 7, 0, 0, 0, 177, 176, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 180,
		1, 0, 0, 0, 179, 181, 7, 1, 0, 0, 180, 179, 1, 0, 0, 0, 181, 182, 1, 0,
		0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 36, 1, 0, 0, 0,
		184, 188, 5, 34, 0, 0, 185, 187, 8, 2, 0, 0, 186, 185, 1, 0, 0, 0, 187,
		190, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191,
		1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 191, 192, 5, 34, 0, 0, 192, 38, 1, 0,
		0, 0, 193, 194, 5, 40, 0, 0, 194, 40, 1, 0, 0, 0, 195, 196, 5, 41, 0, 0,
		196, 42, 1, 0, 0, 0, 197, 198, 5, 123, 0, 0, 198, 44, 1, 0, 0, 0, 199,
		200, 5, 125, 0, 0, 200, 46, 1, 0, 0, 0, 201, 202, 5, 61, 0, 0, 202, 48,
		1, 0, 0, 0, 203, 204, 5, 62, 0, 0, 204, 50, 1, 0, 0, 0, 205, 206, 5, 62,
		0, 0, 206, 207, 5, 61, 0, 0, 207, 52, 1, 0, 0, 0, 208, 209, 5, 60, 0, 0,
		209, 54, 1, 0, 0, 0, 210, 211, 5, 60, 0, 0, 211, 212, 5, 61, 0, 0, 212,
		56, 1, 0, 0, 0, 213, 214, 5, 33, 0, 0, 214, 58, 1, 0, 0, 0, 215, 216, 5,
		61, 0, 0, 216, 217, 5, 61, 0, 0, 217, 60, 1, 0, 0, 0, 218, 219, 5, 33,
		0, 0, 219, 220, 5, 61, 0, 0, 220, 62, 1, 0, 0, 0, 221, 222, 5, 38, 0, 0,
		222, 223, 5, 38, 0, 0, 223, 64, 1, 0, 0, 0, 224, 225, 5, 124, 0, 0, 225,
		226, 5, 124, 0, 0, 226, 66, 1, 0, 0, 0, 227, 228, 5, 43, 0, 0, 228, 68,
		1, 0, 0, 0, 229, 230, 5, 45, 0, 0, 230, 70, 1, 0, 0, 0, 231, 232, 5, 47,
		0, 0, 232, 72, 1, 0, 0, 0, 233, 234, 5, 42, 0, 0, 234, 74, 1, 0, 0, 0,
		235, 236, 5, 58, 0, 0, 236, 76, 1, 0, 0, 0, 237, 241, 7, 3, 0, 0, 238,
		240, 7, 4, 0, 0, 239, 238, 1, 0, 0, 0, 240, 243, 1, 0, 0, 0, 241, 239,
		1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 78, 1, 0, 0, 0, 243, 241, 1, 0,
		0, 0, 244, 245, 5, 59, 0, 0, 245, 80, 1, 0, 0, 0, 246, 247, 5, 46, 0, 0,
		247, 82, 1, 0, 0, 0, 248, 249, 5, 44, 0, 0, 249, 84, 1, 0, 0, 0, 250, 252,
		7, 5, 0, 0, 251, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 251, 1, 0,
		0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 6, 42, 0, 0,
		256, 86, 1, 0, 0, 0, 257, 258, 5, 47, 0, 0, 258, 259, 5, 47, 0, 0, 259,
		263, 1, 0, 0, 0, 260, 262, 8, 6, 0, 0, 261, 260, 1, 0, 0, 0, 262, 265,
		1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 266, 1, 0,
		0, 0, 265, 263, 1, 0, 0, 0, 266, 267, 6, 43, 0, 0, 267, 88, 1, 0, 0, 0,
		7, 0, 177, 182, 188, 241, 253, 263, 1, 6, 0, 0,
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

// GoliteLexerInit initializes any static state used to implement GoliteLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGoliteLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteLexerInit() {
	staticData := &golitelexerLexerStaticData
	staticData.once.Do(golitelexerLexerInit)
}

// NewGoliteLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGoliteLexer(input antlr.CharStream) *GoliteLexer {
	GoliteLexerInit()
	l := new(GoliteLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &golitelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "GoliteLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GoliteLexer tokens.
const (
	GoliteLexerPRINT      = 1
	GoliteLexerSTRUCT     = 2
	GoliteLexerTRUE       = 3
	GoliteLexerFALSE      = 4
	GoliteLexerNIL        = 5
	GoliteLexerNEW        = 6
	GoliteLexerINT        = 7
	GoliteLexerBOOL       = 8
	GoliteLexerSCAN       = 9
	GoliteLexerDELETE     = 10
	GoliteLexerIF         = 11
	GoliteLexerELSE       = 12
	GoliteLexerFOR        = 13
	GoliteLexerRETURN     = 14
	GoliteLexerTYPE       = 15
	GoliteLexerVAR        = 16
	GoliteLexerFUNC       = 17
	GoliteLexerNUMBER     = 18
	GoliteLexerSTRING     = 19
	GoliteLexerLPAREN     = 20
	GoliteLexerRPAREN     = 21
	GoliteLexerLBRACE     = 22
	GoliteLexerRBRACE     = 23
	GoliteLexerEQUALS     = 24
	GoliteLexerGTN        = 25
	GoliteLexerGETN       = 26
	GoliteLexerLTN        = 27
	GoliteLexerLETN       = 28
	GoliteLexerNOT        = 29
	GoliteLexerEQ         = 30
	GoliteLexerNEQ        = 31
	GoliteLexerAND        = 32
	GoliteLexerOR         = 33
	GoliteLexerPLUS       = 34
	GoliteLexerMINUS      = 35
	GoliteLexerFSLASH     = 36
	GoliteLexerASTERISK   = 37
	GoliteLexerCOLON      = 38
	GoliteLexerIDENTIFIER = 39
	GoliteLexerSEMICOLON  = 40
	GoliteLexerDOT        = 41
	GoliteLexerCOMMA      = 42
	GoliteLexerWS         = 43
	GoliteLexerCOMMENT    = 44
)