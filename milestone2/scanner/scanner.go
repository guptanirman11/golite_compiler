package scanner

import (
	"milestone2/token"
)

type Scanner struct {
	input string
	index int
	count int
	// ch        byte
	line   int
	lexeme string
	accept string

	// TODO: IMPLEMENT ME!
}

func New(input string) *Scanner {
	//Feel free to change this implementation
	scanner := &Scanner{input: input, index: 0, line: 1}
	return scanner
}

func (l *Scanner) nextChar() byte {

	var ch byte
	if l.index < len(l.input) {
		ch = l.input[l.index]
		l.index++

	}
	return ch

}

func (l *Scanner) NextToken() token.Token {
	// TODO: Implement Me
	// var tok token.Token
	l.lexeme = ""
	l.accept = ""
	l.count = 0

	// start:
	char := l.nextChar()
	l.lexeme = l.lexeme + string(char)
	l.count++

	switch {
	case isWhitespace(char):
		l.index++
		l.accept = token.SPACE
		goto end
	case char == '\n':
		l.index++
		l.accept = token.NEWLINE
		l.line++
		goto end
	case char == '+':
		// tok = newToken(token.PLUS, l.ch)
		l.index++
		l.accept = token.PLUS
		goto end
	case char == '-':
		l.index++
		l.accept = token.MINUS
		goto end
	case char == '/':
		l.index++
		l.accept = token.DIVIDE
		goto end
	case char == '*':
		l.index++
		l.accept = token.MULTIPLY
		goto end
	case char == '=':
		l.index++
		l.accept = token.EQUAL
		goto end

	case char == ';':
		l.index++
		l.accept = token.SEMICOLON
		goto end

	default:
		if isDigit(char) {
			// l.index++
			goto sint
		} else if isLetter(char) {
			goto sid
		} else {
			goto end
		}

	}
sint:
	char = l.nextChar()
	l.lexeme = l.lexeme + string(char)
	l.count = 1
	l.accept = token.INT_LIT
	if isDigit(char) {
		goto sint
	} else {
		goto end
	}

sid:
	char = l.nextChar()
	// fmt.Println(l.lexeme, string(char))
	l.lexeme = l.lexeme + string(char)
	l.count = 1
	l.accept = token.ID

	if l.lexeme == "let" {
		goto slet
	} else if l.lexeme == "print" {
		goto sprint
	}
	if isDigit(char) || isLetter(char) {
		goto sid
	} else {
		goto end
	}

slet:
	char = l.nextChar()
	l.lexeme = l.lexeme + string(char)
	l.count = 1
	l.accept = token.LET
	if isDigit(char) || isLetter(char) {
		goto sid
	} else {
		goto end
	}

sprint:
	char = l.nextChar()
	l.lexeme = l.lexeme + string(char)
	l.count = 1
	l.accept = token.PRINT
	if isDigit(char) || isLetter(char) {
		goto sid
	} else {
		goto end
	}

end:
	var tok token.Token

	if l.accept != "" {
		// resetting the lexeme

		l.lexeme = l.lexeme[:len(l.lexeme)-1]
		l.index = l.index - l.count
		tok.Type = l.accept
		tok.Literal = l.lexeme
		tok.Line = l.line

	} else if l.index == len(l.input) {
		tok.Type = "EOF"
		tok.Literal = l.lexeme
		tok.Line = l.line
	} else {
		tok.Type = token.ILLEGAL
		tok.Literal = l.lexeme
		tok.Line = l.line
	}

	return tok
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t'
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}
