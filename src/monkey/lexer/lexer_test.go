package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextTokenEqual(t *testing.T) {
	input := `
	10 == 10;
	10 != 9;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
func TestPeekChar(t *testing.T) {
	input := `abc`
	l := New(input)
	pch := l.peekChar()
	if pch != 'b' {
		t.Fatalf("expected b, got %q", pch)
	}
	l.readChar()
	pch = l.peekChar()
	if pch != 'c' {
		t.Fatalf("expected c, got %q", pch)
	}
	l.readChar()
	pch = l.peekChar()
	if pch != 0 {
		t.Fatalf("expected 0, got %q", pch)
	}
}

func TestNextTokenTrueFalseIfElse(t *testing.T) {
	input := `
	if (5 < 10) {
		return true
	} else {
		return false
	};`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}

func TestNextTokenMinusAsterisk(t *testing.T) {
	input := `!-/*5;
	5 < 10 > 5;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenFunction(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	
	let result = add(five, ten);
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestReadNumber(t *testing.T) {
	input := `let x=50;`
	l := New(input)
	idf := l.readIdentifier()
	if idf != "let" {
		t.Fatalf("expected a identifier, got %q", idf)
	}
	l.readChar()
	idf = l.readIdentifier()
	if idf != "x" {
		t.Fatalf("expected a identifier, got %q", idf)
	}
	l.readChar()
	idf = l.readNumber()
	if idf != "50" {
		t.Fatalf("expected a identifier, got %q", idf)
	}

}

func TestIsDigit(t *testing.T) {
	input := `a`
	l := New(input)
	if isDigit(l.ch) {
		t.Fatalf("expected a non digit, got %q", l.ch)
	}

	input = `1`
	l = New(input)
	if !isDigit(l.ch) {
		t.Fatalf("expected a digit, got %q", l.ch)
	}
}

func TestReadIdentifier(t *testing.T) {
	input := `let x=5;`
	l := New(input)
	idf := l.readIdentifier()
	if idf != "let" {
		t.Fatalf("expected a identifier, got %q", idf)
	}
	l.readChar()
	idf = l.readIdentifier()
	if idf != "x" {
		t.Fatalf("expected a identifier, got %q", idf)
	}

	input = `let x_y = 5;`
	l = New(input)
	idf = l.readIdentifier()
	if idf != "let" {
		t.Fatalf("expected a identifier, got %q", idf)
	}
	l.readChar()
	idf = l.readIdentifier()
	if idf != "x_y" {
		t.Fatalf("expected a identifier, got %q", idf)
	}

}

func TestIsLetter(t *testing.T) {
	input := `a`
	l := New(input)
	if !isLetter(l.ch) {
		t.Fatalf("expected a letter, got %q", l.ch)
	}

	input = `+=`
	l = New(input)
	if isLetter(l.ch) {
		t.Fatalf("expected a non letter, got %q", l.ch)
	}

	input = `abc_abc`
	l = New(input)
	for _ = range input {
		if !isLetter(l.ch) {
			t.Fatalf("expected a letter, got %q", l.ch)
		}
		l.readChar()
	}

}

func TestNextTokenWhiteSpaceNewLine(t *testing.T) {
	input := `                   
	
	
	=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
func TestNextTokenWhiteSpace(t *testing.T) {
	input := `                   =+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenIdent(t *testing.T) {
	input := `let five = 5;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestReadChar(t *testing.T) {
	input := `=+(){},;`
	l := New(input)
	if '=' != l.ch {
		t.Fatalf("test failed, expected=%q, got=%q",
			'=', l.ch)
	}

	input = `abcd`
	l = New(input)
	for i := 0; i < len(input); i++ {
		c := input[i]
		if l.ch != c {
			t.Fatalf("test failed, expected=%q, got=%q",
				c, l.ch)
		}
		l.readChar()
	}

}
