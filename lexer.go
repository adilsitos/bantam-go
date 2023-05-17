package main

import "fmt"

type Lexer struct {
	mIndex    int
	mText     string
	mPunctors []byte
}

func NewLexer(text string) *Lexer {
	var mPunctors []byte
	for _, val := range CharacterPunctuator {
		mPunctors = append(mPunctors, val)
	}
	fmt.Println(mPunctors)

	return &Lexer{
		mIndex:    0,
		mText:     text,
		mPunctors: mPunctors,
	}
}

func (l *Lexer) next() *Token {
	for l.mIndex < len(l.mText) {
		ch := l.mText[l.mIndex]
		l.mIndex++
		if l.containPunctor(ch) {
			return NewToken(CharacterTypePunctuactor[ch], string(ch))
		} else if isLetter(ch) {
			start := l.mIndex - 1
			for l.mIndex < len(l.mText) {
				if !isLetter(l.mText[l.mIndex]) {
					break
				}
				l.mIndex++
			}

			name := l.mText[start:l.mIndex]
			fmt.Println(name)
			return NewToken("name", name)
		}
	}

	return NewToken("EOF", "")
}

func (l *Lexer) containPunctor(ch byte) bool {
	for _, val := range l.mPunctors {
		if ch == val {
			return true
		}
	}

	return false
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}
