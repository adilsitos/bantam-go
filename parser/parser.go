package parser

import (
	"log"

	"github.com/adilsitos/bantam-go/expression"
	"github.com/adilsitos/bantam-go/lexer"
	"github.com/adilsitos/bantam-go/token"
)

type Parser struct {
	Tokens          []*token.Token
	Lexer           lexer.Lexer
	PrefixParselets map[token.TokenType]PrefixParselet
}

func NewParser(lexer lexer.Lexer, tokens []*token.Token) *Parser {
	prefixParselet := make(map[token.TokenType]PrefixParselet)

	parser := &Parser{
		Tokens:          tokens,
		Lexer:           lexer,
		PrefixParselets: prefixParselet,
	}

	parser.register(token.NAME, NewNameParselet())
	parser.prefix(token.PLUS)
	parser.prefix(token.MINUS)
	parser.prefix(token.TILDE)
	parser.prefix(token.BANG)

	return parser
}

func (p *Parser) parseExpression() expression.Expression {
	token := p.consume()

	prefix := p.PrefixParselets[token.Type]

	if prefix == nil {
		log.Panicf("could not parse %s", token.Value)
	}

	return prefix.parse(*p, *token)

}

func (p *Parser) consume() *token.Token {
	token := p.lookAhead(0)
	p.removeToken()

	return token
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

func (p *Parser) prefix(tokenType token.TokenType) {
	p.PrefixParselets[tokenType] = NewPrefixParselet()

}
