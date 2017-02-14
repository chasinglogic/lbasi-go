package main

import (
	"fmt"

	"github.com/chasinglogic/lbasi/interpreter"
)

func main() {
	i := interpreter.New("3 + 4")
	ans, err := i.Expr()

	fmt.Printf("%v, %v\n", ans, err)

	i = interpreter.New("10 + 1 + 2 - 3 + 4 + 6 - 15")
	ans, err = i.Expr()

	fmt.Printf("%v, %v\n", ans, err)
}
