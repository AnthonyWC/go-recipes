package strutils

import (
	"strings"
	"unicode"
)

// Returns the string changed with upper case.
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// Returns the string changed with lower case.
func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

// Returns the string changed to upper case for its first letter.
func ToFirstUpper(s string) string {
	if len(s) < 1 {
		return s
	}
	t := strings.Trim(s, " ")
	t = strings.ToLower(t)
	a := []rune(t)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
