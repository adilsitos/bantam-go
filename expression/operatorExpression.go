package expression

import (
	"fmt"
	"strings"

	"github.com/adilsitos/bantam-go/token"
)

type OperatorExpression struct {
	left     Expression
	operator token.TokenType
	right    Expression
}

func NewOperatorExpression(left Expression, operator token.TokenType, right Expression) *OperatorExpression {
	return &OperatorExpression{
		left:     left,
		operator: operator,
		right:    right,
	}
}

func (oe *OperatorExpression) print(builder strings.Builder) {
	builder.WriteString("(")
	oe.left.print(builder)
	builder.WriteString(fmt.Sprintf(" %s ", oe.operator))
	oe.right.print(builder)
	builder.WriteString(")")
}
