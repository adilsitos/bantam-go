package lexer

import (
	"testing"

	"github.com/adilsitos/bantam-go/token"
	"github.com/stretchr/testify/require"
)

func TestLexer(t *testing.T) {
	tCases := []struct {
		text           string
		expectedTokens []token.Token
	}{
		{
			text: "from + offset(time)",
			expectedTokens: []token.Token{
				{Type: token.NAME, Value: "from"},
				{Type: token.PLUS, Value: "+"},
				{Type: token.NAME, Value: "offset"},
				{Type: token.LEFT_PAREN, Value: "("},
				{Type: token.NAME, Value: "time"},
				{Type: token.RIGHT_PAREN, Value: ")"},
			},
		},
	}

	for _, tCase := range tCases {
		lexer := NewLexer(tCase.text)

		count := 0
		for lexer.index < len(lexer.text) {
			tok := lexer.NextToken()

			t.Log(tok)
			t.Log(tCase.expectedTokens[count])

			require.Equal(t, tok.Type, tCase.expectedTokens[count].Type)
			require.Equal(t, tok.Value, tCase.expectedTokens[count].Value)
			count++
		}
	}
}
