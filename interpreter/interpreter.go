package interpreter

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	CurrentToken Token
}

type Interpreter struct {
	Parser Parser
	Err    error
}

func New(expr string) Interpreter {
	i := Interpreter{
		Lexer: NewLexer(expr),
	}

	return i
}

func (p *Parser) getNextToken() Token {
	for p.CurrentChar != 0 {
		if isWhitespace(p.CurrentChar) {
			p.skipWhitespace()
			continue
		}

		if isNumber(p.CurrentChar) {
			return Token{INT, p.term()}
		}

		return Op(p.term())
	}

	return Token{EOF, ""}
}

func (p *Parser) advance() {
	p.POS += 1

	if p.POS >= len(p.Body) {
		p.CurrentChar = 0
		return
	}

	p.CurrentChar = p.Body[p.POS]
}

func (i *Interpreter) error(msgs ...string) error {
	errMsg := fmt.Sprintf("Error parsing input: %s @ %d", string(p.CurrentChar), p.POS)

	if msgs != nil {
		for _, m := range msgs {
			errMsg += "\n" + m
		}
	}

	i.Err = errors.New(errMsg)
	return i.Err
}

func (p *Parser) skipWhitespace() {
	for p.CurrentChar != 0 && strings.Contains(" \n\t", string(p.CurrentChar)) {
		p.advance()
	}
}

func (p *Parser) term() string {
	result := []rune{p.CurrentChar}
	p.advance()

	for p.CurrentChar != 0 && !strings.Contains(" \n", string(p.CurrentChar)) {
		result = append(result, p.CurrentChar)
		p.advance()
	}

	return string(result)
}

func (p *Parser) Eat(t TokenType) error {
	if p.CurrentToken.Type == t {
		p.CurrentToken = p.getNextToken()
		return nil
	}

	return errors.New("Expected token type: " + t.String())
}

func (p *Interpreter) ExprNoErr() int {
	ans, _ := p.Expr()
	return ans
}

func (p *Parser) int() (int, error) {
	token := p.CurrentToken
	err := p.Eat(INT)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(token.Value)
}

func (i *Interpreter) Expr() (int, error) {
	i.Parser.CurrentToken = i.Parser.getNextToken()

	result, err := i.Parser.int()
	if err != nil {
		return 0, i.error(err.Error())
	}

	for isOpToken(i.Parser.CurrentToken.Type) {
		op := i.Parser.CurrentToken
		i.Parser.CurrentToken = i.Parser.getNextToken()

		next, err := i.Parser.int()
		if err != nil {
			return 0, i.error(err.Error())
		}

		switch op.Type {
		case ADD:
			result = result + next
		case SUB:
			result = result - next
		case MUL:
			result = result * next
		case DIV:
			result = result / next
		default:
			return 0, i.error("Invalid operation.")
		}
	}

	return result, nil
}
