package expression

import "strings"

type NameExpression struct {
	name string
}

func NewNameExpression(name string) *NameExpression {
	return &NameExpression{
		name: name,
	}
}

func (ne *NameExpression) getName() string {
	return ne.name
}

func (ne *NameExpression) print(builder strings.Builder) {
	builder.WriteString(ne.name)
}
