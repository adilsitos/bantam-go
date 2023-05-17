package main

type Token struct {
	mType string
	mText string
}

func NewToken(tokenType, txt string) *Token {
	return &Token{
		mType: tokenType,
		mText: txt,
	}
}

// it is necessary to create a map with all types in str
// eg., LEFT_PAREN is a int (), so the str version would be
// its name (LEFT_PAREN)
func (t *Token) GetType() string { return t.mType }
func (t *Token) GetText() string { return t.mText }
