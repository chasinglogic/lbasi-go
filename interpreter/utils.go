package interpreter

import "strings"

func isWhitespace(c rune) bool {
	return strings.Contains(" \n\t", string(c))
}

func isNumber(c rune) bool {
	return strings.Contains("0123456789", string(c))
}
