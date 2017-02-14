package interpreter

import "testing"

func interpTest(expr string, expected int, t *testing.T) {
	i := New(expr)
	answer := i.Expr()

	if answer != expected {
		t.Errorf("Expected %d Got %d", expected, answer)
	}
}

func TestInterpreter(t *testing.T) {
	interpTest("3 + 4", 7, t)
	interpTest("4 - 3", 1, t)
	interpTest("3 / 3", 1, t)
	interpTest("3 * 4", 12, t)
}
