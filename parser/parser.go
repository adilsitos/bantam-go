package parser

import (
	"log"

	"github.com/adilsitos/bantam-go/expression"
	"github.com/adilsitos/bantam-go/lexer"
	"github.com/adilsitos/bantam-go/token"
)

type Parser struct {
	Tokens          []*token.Token
	Lexer           *lexer.Lexer
	PrefixParselets map[token.TokenType]PrefixParselet
	InfixParselet   map[token.TokenType]InfixParselet
}

func NewParser(lexer *lexer.Lexer, tokens []*token.Token) *Parser {
	prefixParselet := make(map[token.TokenType]PrefixParselet)
	infixParselet := make(map[token.TokenType]InfixParselet)

	parser := &Parser{
		Tokens:          tokens,
		Lexer:           lexer,
		PrefixParselets: prefixParselet,
		InfixParselet:   infixParselet,
	}

	parser.register(token.NAME, NewNameParselet())
	parser.prefix(token.PLUS)
	parser.prefix(token.MINUS)
	parser.prefix(token.TILDE)
	parser.prefix(token.BANG)

	parser.registerInfix(token.QUESTION, NewConditionalParselet())

	parser.infixLeft(token.PLUS)
	parser.infixLeft(token.MINUS)
	parser.infixLeft(token.ASTERISK)
	parser.infixLeft(token.SLASH)
	parser.infixLeft(token.CARET)

	return parser
}

func (p *Parser) parseExpression() expression.Expression {
	token := p.consume()

	prefix := p.PrefixParselets[token.Type]

	if prefix == nil {
		log.Panicf("could not parse %s", token.Value)
	}

	left := prefix.parse(p, token)

	token = p.lookAhead(0)
	infix := p.InfixParselet[token.Type]

	if infix == nil {
		return left
	}

	p.consume()

	return infix.parse(p, left, token)

}

func (p *Parser) consume() *token.Token {
	token := p.lookAhead(0)
	p.removeToken()

	return token
}

func (p *Parser) consumeExpected(expected token.TokenType) *token.Token {
	token := p.lookAhead(0)

	if token.Type != expected {
		log.Panicf("Expected token %s and found %s", expected, token.Type)
	}

	return p.consume()
}

func (p *Parser) lookAhead(dist int) *token.Token {
	for dist >= len(p.Tokens) {
		p.Tokens = append(p.Tokens, p.Lexer.NextToken())
	}

	return p.Tokens[0]
}

func (p *Parser) removeToken() {
	p.Tokens = p.Tokens[1:]
}

func (p *Parser) register(tokenType token.TokenType, parselet PrefixParselet) {
	p.PrefixParselets[tokenType] = parselet
}

func (p *Parser) registerInfix(token token.TokenType, parselet InfixParselet) {
	p.InfixParselet[token] = parselet
}

func (p *Parser) prefix(tokenType token.TokenType) {
	p.PrefixParselets[tokenType] = NewPrefixOperatorParselet()
}

func (p *Parser) postfix(tokenType token.TokenType) {
	p.InfixParselet[tokenType] = NewPostfixOperatorParselet()
}

func (p *Parser) infixLeft(tokenType token.TokenType) {
	p.InfixParselet[tokenType] = NewBinaryOperatorParselet()
}
