package expression

import "strings"

type ConditionalExpression struct {
	condition Expression
	thenBlock Expression
	elseBlock Expression
}

func NewConditionalExpression(condition, thenBlock, elseBlock Expression) ConditionalExpression {
	return ConditionalExpression{
		condition: condition,
		thenBlock: thenBlock,
		elseBlock: elseBlock,
	}
}

func (ce ConditionalExpression) print(builder strings.Builder) {
	builder.WriteString("(")
	ce.condition.print(builder)
	builder.WriteString(" ? ")
	ce.thenBlock.print(builder)
	builder.WriteString(" : ")
	ce.elseBlock.print(builder)
	builder.WriteString(")")
}
