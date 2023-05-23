package main

import expression "github.com/adilsitos/bantam-go/expressions"

type PrefixParselets interface {
	parse(parser Parser, token Token) expression.Expression
}

type NameParserlets struct {
}

func (np *NameParserlets) parse(parser Parser, token Token) expression.Expression {
	return expression.NewNameExpression(token.GetText())
}

type PrefixOperatorParselets struct{}

func (po *PrefixOperatorParselets) parse(parser Parser, token Token) expression.Expression {
	operand := parser.parseExpression()
	return expression.NewPrefixExpression(token.GetType(), operand)
}
