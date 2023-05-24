package main

import (
	"log"

	expression "github.com/adilsitos/bantam-go/expressions"
)

type Parser struct {
	tokens          []Token
	prefixParselets map[string]PrefixParselets
	infixParselets  map[string]InfixParselets
}

func NewParser(lexer Lexer) *Parser {
	tok := *lexer.next()

	var tokens []Token
	for tok.GetType() != "EOF" {
		tokens = append(tokens, tok)
		tok = *lexer.next()
	}
	tokens = append(tokens, tok)

	parser := &Parser{
		tokens: tokens,
	}

	parser.registerPrefix("name", &NameParserlets{})
	//parser.register("assign", &AassignParlet{})
	parser.registerInfix("question", &ConditionalParselet{})
	parser.registerPrefix("left_paren", &GroupParselet{})
	//parser.registerInfix("left_paren", &)

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

	left := prefix.parse(*p, tok)

	tok = p.lookAhead()
	infix := p.infixParselets[tok.GetType()]

	if infix == nil {
		return left
	}

	p.consumeToken()

	return infix.parse(*p, left, tok)
}

func (p *Parser) prefix(tokenType string) {
	p.registerPrefix(tokenType, &PrefixOperatorParselets{})
}

func (p *Parser) registerPrefix(tokenType string, prefix PrefixParselets) {
	p.prefixParselets[tokenType] = prefix
}

func (p *Parser) registerInfix(tokenType string, infix InfixParselets) {
	p.infixParselets[tokenType] = infix
}

func (p *Parser) consumeToken() Token {
	tok := p.tokens[0]
	p.tokens = p.tokens[1:]

	return tok
}

func (p *Parser) ConsumeExpected(tokType string) Token {
	tokAhead := p.lookAhead()
	if tokType != tokAhead.GetType() {
		log.Panicf("Expected %s got %s", tokType, tokAhead.GetType())
	}

	return p.consumeToken()
}

func (p *Parser) lookAhead() Token {
	return p.tokens[1]
}

func (p *Parser) matchAndConsume(expected string) bool {
	tok := p.lookAhead()

	if tok.GetType() != expected {
		return false
	}

	p.consumeToken()

	return true
}
