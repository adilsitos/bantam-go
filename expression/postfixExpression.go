package expression

import (
	"fmt"
	"strings"

	"github.com/adilsitos/bantam-go/token"
)

type PostfixExpression struct {
	left     Expression
	operator token.TokenType
}

func NewPostfixExpression(left Expression, operator token.TokenType) *PostfixExpression {
	return &PostfixExpression{
		left:     left,
		operator: operator,
	}
}

func (pe *PostfixExpression) print(builder strings.Builder) {
	builder.WriteString("(")
	pe.left.print(builder)
	builder.WriteString(fmt.Sprintf("%s)", pe.operator))

}
