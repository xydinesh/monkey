package lexer

import (
	"testing"
	"monkey/token"
)

func TestReadNumber(t *testing.T) {
	input := `let x=50;`
	l := New(input)
	idf := l.readIdentifier()
	if  idf != "let" {
		t.Fatalf("expected a identifier, got %q", idf)
	}
	l.readChar()
	idf = l.readIdentifier()
	if  idf != "x" {
		t.Fatalf("expected a identifier, got %q", idf)
	}
	l.readChar()
	idf = l.readNumber()
	if  idf != "50" {
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
	if  idf != "let" {
		t.Fatalf("expected a identifier, got %q", idf)
	}
	l.readChar()
	idf = l.readIdentifier()
	if  idf != "x" {
		t.Fatalf("expected a identifier, got %q", idf)
	}

	input = `let x_y = 5;`
	l = New(input)
	idf = l.readIdentifier()
	if  idf != "let" {
		t.Fatalf("expected a identifier, got %q", idf)
	}
	l.readChar()
	idf = l.readIdentifier()
	if  idf != "x_y" {
		t.Fatalf("expected a identifier, got %q", idf)
	}

}

func TestIsLetter(t *testing.T){
	input := `a`
	l := New(input)
	if  !isLetter(l.ch) {
		t.Fatalf("expected a letter, got %q", l.ch)
	}

	input = `+=`
	l = New(input)
	if  isLetter(l.ch) {
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

func TestNextTokenWhiteSpaceNewLine(t *testing.T){
	input := `                   
	
	
	=+(){},;`
	tests := []struct{
		expectedType token.TokenType
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
func TestNextTokenWhiteSpace(t *testing.T){
	input := `                   =+(){},;`
	tests := []struct{
		expectedType token.TokenType
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

func TestNextTokenIdent(t *testing.T){
	input := `let five = 5;`
	tests := []struct{
		expectedType token.TokenType
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

func TestNextToken(t *testing.T){
	input := `=+(){},;`
	tests := []struct{
		expectedType token.TokenType
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

func TestReadChar(t *testing.T){
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