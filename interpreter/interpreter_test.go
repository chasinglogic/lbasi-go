package interpreter

import "testing"

func interpTest(expr string, expected int, t *testing.T) {
	i := New(expr)
	answer, err := i.Expr()

	if err != nil {
		t.Errorf("Failed with err: %v\n", err)
	}

	if answer != expected {
		t.Errorf("Expected %d Got %d\n", expected, answer)
	}
}

func TestInterpreter(t *testing.T) {
	interpTest("3 + 4", 7, t)
	interpTest("4 - 3", 1, t)
	interpTest("3 / 3", 1, t)
	interpTest("3 * 4", 12, t)
	interpTest("10 + 1 + 2 - 3 + 4 + 6 - 15", 5, t)
	interpTest("7 - 3 + 2 - 1", 5, t)
}
