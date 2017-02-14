package interpreter

import "fmt"

type TokenType int

func (tt TokenType) String() string {
	switch tt {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case DIV:
		return "/"
	case MUL:
		return "*"
	case INT:
		return "INTEGER"
	default:
		return "EOF"
	}
}

const (
	ADD TokenType = iota
	SUB
	MUL
	DIV

	INT

	EOF
)

type Token struct {
	Type  TokenType
	Value string
}

func Op(c string) Token {
	switch c {
	case "+":
		return Token{ADD, "+"}
	case "-":
		return Token{SUB, "-"}
	case "*":
		return Token{MUL, "*"}
	case "/":
		return Token{DIV, "/"}
	default:
		return Token{EOF, ""}
	}
}

func isOpToken(tt TokenType) bool {
	return tt == ADD || tt == SUB || tt == MUL || tt == DIV
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%s, %s)", t.Type.String(), t.Value)
}
