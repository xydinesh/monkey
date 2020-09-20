package ast

import (
	"monkey/token"
)

// Node interface
type Node interface {
	TokenLiteral() string
}

// Statement interface
type Statement interface {
	Node
	statementNode()
}

// Expression interface
type Expression interface {
	Node
	expressionNode()
}

// Program structure
type Program struct {
	Statements []Statement
}

// TokenLiteral implementaion of the program node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement struct
type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression
}

// TokenLiteral function for LetStatement
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// statementNode implementation for Let
func (ls *LetStatement) statementNode() {}

// Identifier type
type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

// TokenLiteral implementation for Identified
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// expressionNode implementation
func (i *Identifier) expressionNode() {}
