package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	// ILLEGAL character special tokens
	ILLEGAL = "ILLEGAL"
	//EOF token
	EOF = "EOF"
	// IDENT for identifier token
	IDENT = "IDENT"
	// INT for integer
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"
	BANG = "!"
	MINUS = "-"
	SLASH = "/"
	ASTERISK = "*"

	EQ = "=="
	NOT_EQ = "!="

	LT = "<"
	GT = ">"

	// Delimeters
	SEMICOLON = ";"
	COMMA = ","

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
	TRUE = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
  }

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}