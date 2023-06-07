package parser

import (
	"github.com/adilsitos/bantam-go/expression"
	"github.com/adilsitos/bantam-go/parser"
	"github.com/adilsitos/bantam-go/token"
)

type PrefixOperatorParselet struct {
}

func (p *PrefixOperatorParselet) parse(parser parser.Parser, token token.Token) expression.Expression {
	operand := parser.parseExpression()

	return PrefixExpression(token.Type, operand)
}
