package lexer

import (
	"monkeypl/token"
	"unicode"
)

type Lexer struct {
	filename     string
	source       string
	position     int
	line         int
	column       int
	nextPosition int

	ch rune
}

func New(file, source string) *Lexer {
	l := &Lexer{filename: file, source: source}
	l.line = 1
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// updates the lines and columns while reading
	if l.ch == '\n' {
		l.line += 1
		l.column = 1
	} else {
		l.column += 1
	}

	if l.nextPosition >= len(l.source) {
		l.ch = 0
	} else {
		l.ch = rune(l.source[l.nextPosition])
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {

	l.eatWhiteSpace()

	var tok token.Token
	switch l.ch {
	case '=':
		tok = l.createToken(string(l.ch), token.ASSIGN)
	case '+':
		tok = l.createToken(string(l.ch), token.PLUS)
	case '-':
		tok = l.createToken(string(l.ch), token.MINUS)
	case '!':
		tok = l.createToken(string(l.ch), token.BANG)
	case '<':
		tok = l.createToken(string(l.ch), token.LT)
	case '>':
		tok = l.createToken(string(l.ch), token.GT)
	case '*':
		tok = l.createToken(string(l.ch), token.ASTERISK)
	case '/':
		tok = l.createToken(string(l.ch), token.SLASH)
	case '{':
		tok = l.createToken(string(l.ch), token.LBRACE)
	case '}':
		tok = l.createToken(string(l.ch), token.RBRACE)
	case '(':
		tok = l.createToken(string(l.ch), token.LPAREN)
	case ')':
		tok = l.createToken(string(l.ch), token.RPAREN)
	case ',':
		tok = l.createToken(string(l.ch), token.COMMA)
	case ';':
		tok = l.createToken(string(l.ch), token.SEMICOLON)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if unicode.IsLetter(l.ch) {
			literal := l.readIdentifier()
			tType := token.LookupIdentifier(literal)
			tok = l.createToken(literal, tType)
			return tok
		} else if unicode.IsDigit(l.ch) {
			num, tType := l.readDigit()
			tok = l.createToken(num, tType)
			return tok
		} else {
			tok = l.createToken(string(l.ch), token.ILLEGAL)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) eatWhiteSpace() {
	for l.ch == '\t' || l.ch == ' ' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) readDigit() (string, token.TokenType) {
	hasDecimal := false
	position := l.position
	var tType token.TokenType = token.INT

	for unicode.IsDigit(l.ch) || l.ch == '.' {
		if l.ch == '.' {
			if hasDecimal {
				tType = token.ILLEGAL
			} else {
				hasDecimal = true
				tType = token.FLOAT
			}
		}
		l.readChar()
	}
	return l.source[position:l.position], tType
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for unicode.IsLetter(l.ch) {
		l.readChar()
	}
	return l.source[start:l.position]
}

func (l *Lexer) createToken(literal string, tokenType token.TokenType) token.Token {
	return token.Token{
		Type:     tokenType,
		Literal:  literal,
		FileName: l.filename,
		Line:     l.line,
		Column:   l.column,
	}
}
