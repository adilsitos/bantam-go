package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLexer(t *testing.T) {
	tCases := []struct {
		input             string
		expectedTokenType []string
		expectedText      []string
	}{
		{
			input:             "from + offset(time)",
			expectedTokenType: []string{"name", "plus", "name", "left_paren", "name", "right_paren", "EOF"},
			expectedText:      []string{"from", "+", "offset", "(", "time", ")", ""},
		},
	}

	for _, tCase := range tCases {
		lexer := NewLexer(tCase.input)

		var tokens []*Token

		stop := false
		for !stop {
			tok := lexer.next()
			tokens = append(tokens, tok)
			t.Log(tok)

			if tok.GetType() == "EOF" {
				stop = true
			}
		}

		for i, tok := range tokens {
			require.Equal(t, tCase.expectedTokenType[i], tok.GetType())
			require.Equal(t, tCase.expectedText[i], tok.GetText())
		}
	}

}
