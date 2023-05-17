package main

const (
	LEFT_PAREN = iota
	RIGHT_PAREN
	COMMA
	ASSIGN
	PLUS
	MINUS
	ASTERISK
	SLASH
	CARET
	TILDE
	BANG
	QUESTION
	COLON
	NAME
	EOF
)

var CharacterPunctuator map[int]byte = map[int]byte{
	LEFT_PAREN:  '(',
	RIGHT_PAREN: ')',
	COMMA:       ',',
	ASSIGN:      '=',
	PLUS:        '+',
	MINUS:       '-',
	ASTERISK:    '*',
	SLASH:       '/',
	CARET:       '^',
	TILDE:       '~',
	BANG:        '!',
	QUESTION:    '?',
	COLON:       ':',
}

var CharacterTypePunctuactor map[byte]string = map[byte]string{
	'(': "left_paren",
	')': "right_paren",
	',': "comma",
	'=': "assign",
	'+': "plus",
	'-': "minus",
	'*': "asterisk",
	'/': "slash",
	'^': "caret",
	'~': "tilde",
	'!': "bang",
	'?': "question",
	':': "colon",
}
