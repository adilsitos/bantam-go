package main

import (
	"log"

	expression "github.com/adilsitos/bantam-go/expressions"
)

type Parser struct {
	tokens          []Token
	prefixParselets map[string]PrefixParselets
}

func NewParser(tokens []Token) *Parser {
	parser := &Parser{
		tokens: tokens,
	}

	parser.register("name", &NameParserlets{})
	parser.prefix("plus")
	parser.prefix("minus")
	parser.prefix("tilde")
	parser.prefix("bang")

	return parser
}

func (p *Parser) parseExpression() expression.Expression {
	tok := p.consumeToken()
	prefix := p.prefixParselets[tok.GetType()]

	if prefix == nil {
		log.Panicf("could not find prefix for %s", tok.GetType())
	}

	return prefix.parse(*p, tok)
}

func (p *Parser) prefix(tokenType string) {
	p.register(tokenType, &PrefixOperatorParselets{})
}

func (p *Parser) register(tokenType string, prefix PrefixParselets) {
	p.prefixParselets[tokenType] = prefix
}

func (p *Parser) consumeToken() Token {
	tok := p.tokens[0]
	p.tokens = p.tokens[1:]

	return tok
}
