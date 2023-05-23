package expression

import (
	"fmt"
	"strings"
)

type PostfixExpression struct {
	left     Expression
	operator string
}

func NewPostfixExpression(left Expression, operator string) PostfixExpression {
	return PostfixExpression{
		left:     left,
		operator: operator,
	}
}

func (pe PostfixExpression) print(builder strings.Builder) {
	builder.WriteString("(")
	pe.left.print(builder)
	builder.WriteString(fmt.Sprintf("%s)", pe.operator))
}
