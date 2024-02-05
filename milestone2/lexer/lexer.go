package lexer

import (
	// "cal/context"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Lexer interface {
	GetTokenStream() *antlr.CommonTokenStream
	PrintTokens()
	// GetErrors() []*context.CompilerError
}

type lexerWrapper struct {
	// *antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	// errors                      []*context.CompilerError
	antlrLexer *GoliteLexer
	stream     *antlr.CommonTokenStream
}

func (lexer *lexerWrapper) GetTokenStream() *antlr.CommonTokenStream {
	return lexer.stream
}
func (lex *lexerWrapper) PrintTokens() {
	lex.stream.Fill()
	for _, token := range lex.stream.GetAllTokens() {
		fmt.Println(token.GetText())
	}
}
func (lex *lexerWrapper) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// lex.errors = append(lex.errors, &context.CompilerError{
	// 	Line:   line,
	// 	Column: column,
	// 	Msg:    msg,
	// 	Phase:  context.LEXER,
	// })
}

//	func (lexer *lexerWrapper) GetErrors() []*context.CompilerError {
//		return lexer.errors
//	}
func NewLexer(inputSourcePath string) Lexer {
	input, _ := antlr.NewFileStream(inputSourcePath)
	// lexer := &lexerWrapper{antlr.NewDefaultErrorListener(), nil, nil, nil}
	antlrLexer := NewGoliteLexer(input)
	lexer := &lexerWrapper{antlrLexer, antlr.NewCommonTokenStream(antlrLexer, 0)}
	// antlrLexer.RemoveErrorListeners()
	// antlrLexer.AddErrorListener(lexer)
	// tokenStream := antlr.NewCommonTokenStream(antlrLexer, 0)
	// lexer.antrlLexer = antlrLexer
	// lexer.stream = tokenStream
	return lexer

}
