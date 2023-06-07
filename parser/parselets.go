package parser

import (
	"github.com/adilsitos/bantam-go/expression"
	"github.com/adilsitos/bantam-go/token"
)

type PrefixParselet interface {
	parse(parser Parser, token token.Token) expression.Expression
}

type NameParselet struct{}

func NewNameParselet() NameParselet {
	return NameParselet{}
}

func (np NameParselet) parse(parser Parser, token token.Token) expression.Expression {
	return expression.NewNameExpression(token.Value)
}

type PrefixOperator struct{}

func NewPrefixOperator() PrefixOperator {
	return PrefixOperator{}
}

func (po PrefixOperator) parse(parser Parser, token token.Token) expression.Expression {
	right := parser.parseExpression()

	return expression.NewPrefixExpression(token.Type, right)
}
