package expression

import (
	"fmt"
	"strings"
)

type PrefixExpression struct {
	operator string
	right    Expression
}

func NewPrefixExpression(operator string, right Expression) *PrefixExpression {
	return &PrefixExpression{
		operator: operator,
		right:    right,
	}
}

func (pe *PrefixExpression) print(builder strings.Builder) {
	builder.WriteString(fmt.Sprintf("(%s", pe.operator))
	pe.right.print(builder)
	builder.WriteString(")")

	fmt.Println(builder.String())
}
