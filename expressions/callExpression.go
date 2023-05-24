package expression

import (
	"strings"
)

type CallExpression struct {
	function Expression
	args     []Expression
}

func NewCallExpression(fn Expression, args []Expression) CallExpression {
	return CallExpression{
		function: fn,
		args:     args,
	}
}

func (ce CallExpression) print(builder strings.Builder) {
	ce.function.print(builder)

	builder.WriteString("(")
	for i := 0; i < len(ce.args); i++ {
		ce.args[0].print(builder)
		ce.args = ce.args[1:]
		if i < len(ce.args)-1 {
			builder.WriteString(", ")
		}
	}

	builder.WriteString(")")
}
