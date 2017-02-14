package interpreter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Interpreter struct {
	Body         []rune
	POS          int
	CurrentToken Token
	CurrentChar  *rune
}

func New(expr string) Interpreter {
	return Interpreter{
		Body: []rune(expr),
		POS:  0,
	}
}

func (i *Interpreter) getNextToken() Token {
	for i.CurrentChar != nil {
		if isWhitespace(*i.CurrentChar) {
			i.skipWhitespace()
			continue
		}

		if isNumber(*i.CurrentChar) {
			return Token{INT, i.term()}
		}

		return Op(*i.CurrentChar)
	}

	return Token{EOF, ""}
}

func (i *Interpreter) advance() {
	i.POS += 1

	if i.POS > len(i.Body) {
		i.CurrentChar = nil
	}

	i.CurrentChar = &i.Body[i.POS]
}

func (i *Interpreter) error() {
	fmt.Println("Error parsing input character: %v @ %d", i.CurrentChar, i.POS)
	os.Exit(1)
}

func (i *Interpreter) skipWhitespace() {
	for i.CurrentChar != nil && !strings.Contains(" \n\t", string(*i.CurrentChar)) {
		i.advance()
	}
}

func (i *Interpreter) term() string {
	result := []rune{}

	for i.CurrentChar != nil && !strings.Contains(" \n", string(*i.CurrentChar)) {
		result = append(result, *i.CurrentChar)
	}

	return string(result)
}

func (i *Interpreter) Eat(t TokenType) {
	if i.CurrentToken.Type == t {
		i.CurrentToken = i.getNextToken()
		return
	}

	i.error()
}

func (i *Interpreter) Expr() int {
	i.CurrentToken = i.getNextToken()

	left := i.CurrentToken
	i.Eat(INT)

	op := i.CurrentToken
	switch op.Type {
	case ADD:
		i.Eat(ADD)
	case SUB:
		i.Eat(SUB)
	case MUL:
		i.Eat(MUL)
	case DIV:
		i.Eat(DIV)
	default:
		i.error()
	}

	right := i.CurrentToken
	i.Eat(INT)

	l, _ := strconv.Atoi(left.Value)
	r, _ := strconv.Atoi(right.Value)

	switch op.Type {
	case ADD:
		return l + r
	case SUB:
		return l - r
	case MUL:
		return l * r
	case DIV:
		return l / r
	default:
		return 0
	}

}
