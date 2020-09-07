package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"

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
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
  }

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}