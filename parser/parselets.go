package parser

import (
	"github.com/adilsitos/bantam-go/expression"
	"github.com/adilsitos/bantam-go/token"
)

type PrefixParselet interface {
	parse(parser *Parser, token *token.Token) expression.Expression
}

type NameParselet struct{}

func NewNameParselet() NameParselet {
	return NameParselet{}
}

func (np NameParselet) parse(parser *Parser, token *token.Token) expression.Expression {
	return expression.NewNameExpression(token.Value)
}

type PrefixOperatorParselet struct{}

func NewPrefixOperatorParselet() PrefixOperatorParselet {
	return PrefixOperatorParselet{}
}

func (po PrefixOperatorParselet) parse(parser *Parser, token *token.Token) expression.Expression {
	right := parser.parseExpression()

	return expression.NewPrefixExpression(token.Type, right)
}

type InfixParselet interface {
	parse(parser *Parser, left expression.Expression, token *token.Token) expression.Expression
}

type BinaryOperatorParselet struct{}

func NewBinaryOperatorParselet() *BinaryOperatorParselet {
	return &BinaryOperatorParselet{}
}

func (bo *BinaryOperatorParselet) parse(parser *Parser, left expression.Expression, token *token.Token) expression.Expression {
	right := parser.parseExpression()

	return expression.NewOperatorExpression(left, token.Type, right)
}

type PostfixOperatorParselet struct{}

func NewPostfixOperatorParselet() *PostfixOperatorParselet {
	return &PostfixOperatorParselet{}
}

func (po *PostfixOperatorParselet) parse(parser *Parser, left expression.Expression, token *token.Token) expression.Expression {
	return expression.NewPostfixExpression(left, token.Type)
}

type ConditionalParselet struct{}

func NewConditionalParselet() *ConditionalParselet {
	return &ConditionalParselet{}
}

func (co *ConditionalParselet) parse(parser *Parser, left expression.Expression, tokenA *token.Token) expression.Expression {
	thenArm := parser.parseExpression()
	parser.consumeExpected(token.COLON)
	elseArm := parser.parseExpression()

	return expression.NewConditionalExpression(left, thenArm, elseArm)
}
