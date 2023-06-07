package expression

import "github.com/adilsitos/bantam-go/token"

type PrefixExpression struct {
	operator token.TokenType
	right    Expression
}

func NewPrefixExpression(operator token.TokenType, right Expression) *PrefixExpression {
	return &PrefixExpression{
		operator: operator,
		right:    right,
	}
}
