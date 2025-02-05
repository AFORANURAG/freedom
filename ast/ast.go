package ast

import "freedom/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// program is also a node
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// identier produce or resolves to a value, so it an expression and it also a node
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      //
	// basically let for token and identifier name in place of value so let i mai let
}

// let x
// so x is the value of identfier here and its type in let statement is of IDENT since it identifier.
// let name, name is indentifier value and its type is IDENT

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
