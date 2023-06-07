package lexer

import "github.com/adilsitos/bantam-go/token"

type Lexer struct {
	text  string
	index int
}

func NewLexer(text string) *Lexer {
	return &Lexer{
		text:  text,
		index: 0,
	}
}

func (l *Lexer) NextToken() *token.Token {

	for l.index < len(l.text) {
		curr := l.text[l.index]
		l.index++

		if val, ok := token.MapTokenType[curr]; ok {
			return &token.Token{Type: val, Value: string(val)}
		}

		if isLetter(curr) {
			start := l.index - 1

			for l.index < len(l.text) {
				if !isLetter(l.text[l.index]) {
					break
				}
				l.index++
			}

			name := l.text[start:l.index]
			return &token.Token{Type: token.NAME, Value: name}
		}
	}

	return &token.Token{Type: token.EOF, Value: ""}
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}
