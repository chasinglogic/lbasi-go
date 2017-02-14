package interpreter

type TokenType int

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

func Op(c rune) Token {
	switch c {
	case '+':
		return Token{ADD, "+"}
	case '-':
		return Token{SUB, "-"}
	case '*':
		return Token{MUL, "*"}
	case '/':
		return Token{DIV, "/"}
	default:
		return Token{EOF, ""}
	}
}
