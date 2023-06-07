package parselets

import (
	"github.com/adilsitos/bantam-go/expression"
	"github.com/adilsitos/bantam-go/parser"
	"github.com/adilsitos/bantam-go/token"
)

type NameParselet struct{}

func NewNameParselet() *NameParselet {
	return &NameParselet{}
}

func parse(parser parser.Parser, token token.Token) expression.Expression {
	return expression.NewNameExpression(token.Value)
}
