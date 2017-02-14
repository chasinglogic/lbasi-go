package interpreter

type Lexer struct {
	Body        []rune
	POS         int
	CurrentChar rune
}

func NewLexer(expr string) Lexer {
	l := Lexer{
		Body: []rune(expr),
		POS:  0,
	}

	l.CurrentChar = l.Body[0]
	return l
}
