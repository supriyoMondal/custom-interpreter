package ast

import "custom-interpreter/token"

type Node interface {
	TokeLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}
type Program struct {
	Statement []Statement
}

func (p *Program) TokeLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokeLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  Identifier
	Value string
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()     {}
func (i *Identifier) TokeLiteral() string { return i.Token.Literal }
