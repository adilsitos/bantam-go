package expression

import (
	"fmt"
	"strings"
)

type OperatorExpression struct {
	left     Expression
	operator string
	right    Expression
}

func NewOperatorExpression(left Expression, operator string, right Expression) OperatorExpression {
	return OperatorExpression{
		left:     left,
		operator: operator,
		right:    right,
	}
}

func (op OperatorExpression) print(builder strings.Builder) {
	builder.WriteString("(")
	op.left.print(builder)
	builder.WriteString(fmt.Sprintf(" %s ", op.operator))
	op.right.print(builder)
	builder.WriteString(")")
}
