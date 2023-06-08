package expression

import "strings"

type ConditionalExpression struct {
	condition Expression
	thenArm   Expression
	elseArm   Expression
}

func NewConditionalExpression(condition, thenArm, elseArm Expression) *ConditionalExpression {
	return &ConditionalExpression{
		condition: condition,
		elseArm:   elseArm,
		thenArm:   thenArm,
	}
}

func (ce *ConditionalExpression) print(builder strings.Builder) {
	builder.WriteString("(")
	ce.condition.print(builder)
	builder.WriteString(" ? ")
	ce.thenArm.print(builder)
	builder.WriteString(" : ")
	ce.elseArm.print(builder)
	builder.WriteString(")")
}
