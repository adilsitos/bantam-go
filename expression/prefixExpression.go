package expression

import (
	"fmt"
	"strings"

	"github.com/adilsitos/bantam-go/token"
)

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

func (pe *PrefixExpression) print(builder strings.Builder) {
	builder.WriteString(fmt.Sprintf("(%s", pe.operator))
	pe.right.print(builder)
	builder.WriteString(")")
}
