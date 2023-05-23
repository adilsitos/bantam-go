package expression

import (
	"fmt"
	"strings"
)

type NameExpression struct {
	name string
}

func NewNameExpression(name string) *NameExpression {
	return &NameExpression{
		name: name,
	}
}

func (ne *NameExpression) print(builder strings.Builder) {
	fmt.Println(ne.name)
}
