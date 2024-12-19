package lexer

import (
	"monkeypl/token"
	"testing"
)

type nextTokenSubtest struct {
	name  string
	input string
	wants []token.Token
}

func (st nextTokenSubtest) run(t *testing.T) {
	l := New(st.name, st.input)
	for i, tt := range st.wants {
		tok := l.NextToken()
		if tok.Type != tt.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected %q, got %q", i+1, tt.Type, tok.Type)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("tests[%d] - token literal wrong. expected %q, got %q", i+1, tt.Literal, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	subtests := []nextTokenSubtest{
		{
			"simple input", "+={}(),;", []token.Token{
				{Type: token.PLUS, Literal: "+"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			"basic monkey language", "let five = 5;\nlet ten = 10;\nlet add = fn(x, y) {\nx + y;\n};\nlet result = add(five, ten);\n",
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "five"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "ten"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "10"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "add"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.FUNCTION, Literal: "fn"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.IDENT, Literal: "y"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.IDENT, Literal: "y"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "result"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.IDENT, Literal: "add"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "five"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.IDENT, Literal: "ten"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			"basic monkey language (with !-/*)", `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (five < ten) {
	return true;
} elif (five == ten) {
	return 10;
} else {
return false;
}`,
			[]token.Token{
				{Type: token.LET, Literal: "let"}, {Type: token.IDENT, Literal: "five"}, {Type: token.ASSIGN, Literal: "="}, {Type: token.INT, Literal: "5"}, {Type: token.SEMICOLON, Literal: ";"}, {Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "ten"}, {Type: token.ASSIGN, Literal: "="}, {Type: token.INT, Literal: "10"},
				{Type: token.SEMICOLON, Literal: ";"}, {Type: token.LET, Literal: "let"}, {Type: token.IDENT, Literal: "add"},
				{Type: token.ASSIGN, Literal: "="}, {Type: token.FUNCTION, Literal: "fn"}, {Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "x"}, {Type: token.COMMA, Literal: ","}, {Type: token.IDENT, Literal: "y"},
				{Type: token.RPAREN, Literal: ")"}, {Type: token.LBRACE, Literal: "{"}, {Type: token.IDENT, Literal: "x"},
				{Type: token.PLUS, Literal: "+"}, {Type: token.IDENT, Literal: "y"}, {Type: token.SEMICOLON, Literal: ";"},
				{Type: token.RBRACE, Literal: "}"}, {Type: token.SEMICOLON, Literal: ";"}, {Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "result"}, {Type: token.ASSIGN, Literal: "="}, {Type: token.IDENT, Literal: "add"},
				{Type: token.LPAREN, Literal: "("}, {Type: token.IDENT, Literal: "five"}, {Type: token.COMMA, Literal: ","},
				{Type: token.IDENT, Literal: "ten"}, {Type: token.RPAREN, Literal: ")"}, {Type: token.SEMICOLON, Literal: ";"},
				{Type: token.BANG, Literal: "!"}, {Type: token.MINUS, Literal: "-"}, {Type: token.SLASH, Literal: "/"},
				{Type: token.ASTERISK, Literal: "*"}, {Type: token.INT, Literal: "5"}, {Type: token.SEMICOLON, Literal: ";"},
				{Type: token.INT, Literal: "5"}, {Type: token.LT, Literal: "<"}, {Type: token.INT, Literal: "10"},
				{Type: token.GT, Literal: ">"}, {Type: token.INT, Literal: "5"}, {Type: token.SEMICOLON, Literal: ";"},
				{Type: token.IF, Literal: "if"}, {Type: token.LPAREN, Literal: "("}, {Type: token.IDENT, Literal: "five"},
				{Type: token.LT, Literal: "<"}, {Type: token.IDENT, Literal: "ten"}, {Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"}, {Type: token.RETURN, Literal: "return"}, {Type: token.TRUE, Literal: "true"}, {Type: token.SEMICOLON, Literal: ";"},
				{Type: token.RBRACE, Literal: "}"}, {Type: token.ELIF, Literal: "elif"}, {Type: token.LPAREN, Literal: "("}, {Type: token.IDENT, Literal: "five"},
				{Type: token.ASSIGN, Literal: "="}, {Type: token.ASSIGN, Literal: "="}, {Type: token.IDENT, Literal: "ten"}, {Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"}, {Type: token.RETURN, Literal: "return"}, {Type: token.INT, Literal: "10"}, {Type: token.SEMICOLON, Literal: ";"}, {Type: token.RBRACE, Literal: "}"},
				{Type: token.ELSE, Literal: "else"}, {Type: token.LBRACE, Literal: "{"}, {Type: token.RETURN, Literal: "return"}, {Type: token.FALSE, Literal: "false"},
				{Type: token.SEMICOLON, Literal: ";"}, {Type: token.RBRACE, Literal: "}"}, {Type: token.EOF, Literal: ""},
			},
		},
	}

	for _, st := range subtests {
		t.Run(st.name, st.run)
	}
}
