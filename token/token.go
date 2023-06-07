package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

// Token types
const (
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	COMMA       = ","
	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	ASTERISK    = "*"
	SLASH       = "/"
	CARET       = "^"
	TILDE       = "~"
	BANG        = "!"
	QUESTION    = "?"
	COLON       = ":"
	NAME        = "Name"
	EOF         = "EOF"
)

var MapTokenType map[byte]TokenType = map[byte]TokenType{
	'(': LEFT_PAREN,
	')': RIGHT_PAREN,
	',': COMMA,
	'=': ASSIGN,
	'+': PLUS,
	'-': MINUS,
	'*': ASTERISK,
	'/': SLASH,
	'^': CARET,
	'~': TILDE,
	'!': BANG,
	'?': QUESTION,
	':': COLON,
}
