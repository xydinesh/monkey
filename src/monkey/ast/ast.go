package ast

import (
	"bytes"
	"monkey/token"
)

// Node interface
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

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

// String implementation for Identifier
func (i *Identifier) String() string {
	return i.Value
}

// ReturnStatement struct
type ReturnStatement struct {
	Token       token.Token // the return token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral implementation of ReturnStatements
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String implementation of ReturnStatemens
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement struct
type ExpressionStatement struct {
	Token      token.Token // first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral implementation for ExpressionStatements
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String implementation for ExpressionStatements
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// IntegerLiteral type
type IntegerLiteral struct {
	Token token.Token // token.INT token
	Value int64
}

// TokenLiteral implementation for Identified
func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

// expressionNode implementation
func (i *IntegerLiteral) expressionNode() {}

// String implementation for IntegerLiteral
func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

// PrefixExpression type
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// TokenLiteral implementation for PrefixExpression
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String implementation for PrefixExpression
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

func (pe *PrefixExpression) expressionNode() {}

// InfixExpression type
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

// TokenLiteral implementation for InfixExpression
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String implementation for InfixExpression
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

func (ie *InfixExpression) expressionNode() {}

// Boolean type
type Boolean struct {
	Token token.Token
	Value bool
}

// TokenLiteral implementation for Boolean
func (be *Boolean) TokenLiteral() string {
	return be.Token.Literal
}

// String implementation for Boolean
func (be *Boolean) String() string {
	return be.Token.Literal
}

func (ie *Boolean) expressionNode() {}
