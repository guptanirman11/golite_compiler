package parser

import (
	"golite/lexer"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Parser interface {
	// Parse() *ast.Program
	// GetErrors() []*context.CompilerError
	Parse() bool
}

type parserWrapper struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	*BaseGoliteParserListener
	antlrParser *GoliteParser
	lexer       lexer.Lexer
	// errors      []*context.CompilerError
	// nodes       map[string]interface{}
}

// func (parser *parserWrapper) GetErrors() []*context.CompilerError {
// 	return parser.errors
// }

// func (parser *parserWrapper) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
// 	parser.errors = append(parser.errors, &context.CompilerError{
// 		Line:   line,
// 		Column: column,
// 		Msg:    msg,
// 		Phase:  context.PARSER,
// 	})
// }

func NewParser(lexer lexer.Lexer) Parser {
	// parser := &parserWrapper{antlr.NewDefaultErrorListener(), &BaseGoliteParserListener{}, nil, lexer, nil, make(map[string]interface{})}
	// antlrParser := NewGoliteParser(lexer.GetTokenStream())
	// antlrParser.RemoveErrorListeners()
	// antlrParser.AddErrorListener(parser)
	// parser.antlrParser = antlrParser
	// return parser
	parser := &parserWrapper{antlr.NewDefaultErrorListener(), &BaseGoliteParserListener{}, nil, lexer, nil}
	antlrParser := NewGoliteParser(lexer.GetTokenStream())
	parser.antlrParser = antlrParser
	parser.lexer = lexer
	// parser := &parserWrapper{antlrParser, lexer}
	return parser
}
func (parser *parserWrapper) Parse() bool {
	// ctx := parser.antlrParser.Program()
	// if context.HasErrors(parser.lexer.GetErrors()) ||
	// 	context.HasErrors(parser.GetErrors()) {
	// 	return nil
	// }
	// antlr.ParseTreeWalkerDefault.Walk(parser, ctx)
	// _, _, key := GetTokenInfo(ctx)
	// return parser.nodes[key].(*ast.Program)

	if parser.antlrParser.Program() != nil {
		return true
	} else {
		return false
	}
}

/******************* Implementation of the Listeners **************************/

// // GetTokenInfo gerates a unique key for each node in the ParserTree
// func GetTokenInfo(ctx antlr.ParserRuleContext) (int, int, string) {
// 	key := fmt.Sprintf("%d,%d", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
// 	return ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), key
// }
// func (parser *parserWrapper) ExitProgram(c *ProgramContext) {
// 	line, col, key := GetTokenInfo(c)
// 	var stmts []ast.Statement
// 	for _, stmtCtx := range c.AllStatement() {
// 		_, _, key := GetTokenInfo(stmtCtx)
// 		astStmt := parser.nodes[key].(ast.Statement)
// 		stmts = append(stmts, astStmt)
// 	}
// 	parser.nodes[key] = ast.NewProgram(stmts, token.NewToken(line, col))
// }
// func (parser *parserWrapper) ExitDecl(c *DeclContext) {

// 	line, col, key := GetTokenInfo(c)
// 	var expr ast.Expression
// 	if exprCtx := c.GetExpr(); exprCtx != nil {
// 		_, _, exprKey := GetTokenInfo(exprCtx)
// 		expr = parser.nodes[exprKey].(ast.Expression)
// 	} else {
// 		expr = nil
// 	}
// 	decl := ast.NewDeclStm(c.GetTyDec().GetId().GetText(), types.StringToType(c.GetTyDec().GetTy().GetText()), expr, token.NewToken(line, col))
// 	parser.nodes[key] = decl
// }

// // EnterAssignment is called when entering the assignment production.
// func (parser *parserWrapper) ExitAssignment(c *AssignmentContext) {
// 	line, col, key := GetTokenInfo(c)
// 	_, _, eprKey := GetTokenInfo(c.Expression())
// 	expr := parser.nodes[eprKey].(ast.Expression)
// 	parser.nodes[key] = ast.NewAssignStm(c.IDENTIFIER().GetText(), expr, token.NewToken(line, col))
// }

// // EnterPrint is called when entering the print production.
// func (parser *parserWrapper) ExitPrint(c *PrintContext) {
// 	line, col, key := GetTokenInfo(c)
// 	_, _, eprKey := GetTokenInfo(c.Expression())
// 	expr := parser.nodes[eprKey].(ast.Expression)
// 	parser.nodes[key] = ast.NewPrintStm(expr, token.NewToken(line, col))
// }

// // ExitExpression is called when exiting the expression production.
// func (parser *parserWrapper) ExitExpression(c *ExpressionContext) {
// 	_, _, key := GetTokenInfo(c)
// 	termContexts := c.AllExpressionPrime()
// 	_, _, eprKey := GetTokenInfo(c.RelationTerm())
// 	if len(termContexts) == 0 {
// 		expr := parser.nodes[eprKey].(ast.Expression)
// 		parser.nodes[key] = expr
// 	} else {
// 		// Exercise for you: You need to think about how you can build a nested BinOp ast from the termContexts.
// 	}
// }
// func (parser *parserWrapper) ExitRelationTerm(c *RelationTermContext) {

// 	_, _, key := GetTokenInfo(c)
// 	termContexts := c.AllRelationTermPrime()
// 	_, _, eprKey := GetTokenInfo(c.Factor())
// 	if len(termContexts) == 0 {
// 		expr := parser.nodes[eprKey].(ast.Expression)
// 		parser.nodes[key] = expr
// 	} else {
// 		// Exercise for you: You need to think about how you can build a nested BinOp ast from the termContexts.
// 	}
// }
// func (parser *parserWrapper) ExitFactor(c *FactorContext) {
// 	line, col, key := GetTokenInfo(c)

// 	if intFactor := c.NUMBER(); intFactor != nil {
// 		intValue, _ := strconv.Atoi(intFactor.GetText())
// 		parser.nodes[key] = ast.NewIntLit(int64(intValue), token.NewToken(line, col))
// 	} else if trueFactor := c.TRUE(); trueFactor != nil {
// 		parser.nodes[key] = ast.NewBoolLit(true, token.NewToken(line, col))
// 	} else if falseFactor := c.FALSE(); falseFactor != nil {
// 		parser.nodes[key] = ast.NewBoolLit(false, token.NewToken(line, col))
// 	} else if idFactor := c.IDENTIFIER(); idFactor != nil {
// 		parser.nodes[key] = ast.NewVariable(idFactor.GetText(), token.NewToken(line, col))
// 	}
// }
