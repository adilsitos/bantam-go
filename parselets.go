package main

import expression "github.com/adilsitos/bantam-go/expressions"

type PrefixParselets interface {
	parse(parser Parser, token Token) expression.Expression
}

type InfixParselets interface {
	parse(parser Parser, left expression.Expression, token Token) expression.Expression
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

type BinaryOperatorParselets struct{}

func (ip *BinaryOperatorParselets) parse(parser Parser, left expression.Expression, token Token) expression.Expression {
	right := parser.parseExpression()

	return expression.NewOperatorExpression(left, token.GetType(), right)
}

type PostfixOperatorParselets struct{}

func (pp *PostfixOperatorParselets) parse(parser Parser, left expression.Expression, tok Token) expression.Expression {
	return expression.NewPostfixExpression(left, tok.GetType())
}

type ConditionalParselet struct{}

func (co *ConditionalParselet) parse(parser Parser, left expression.Expression, token Token) expression.Expression {
	thenBlock := parser.parseExpression()
	parser.ConsumeExpected(":")
	elseBlock := parser.parseExpression()

	return expression.NewConditionalExpression(left, thenBlock, elseBlock)
}
