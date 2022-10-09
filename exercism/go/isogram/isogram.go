package isogram

import (
	"unicode"
)

var exists = struct{}{}

// IsIsogram returns True when a word is an isogram and False otherwise.
func IsIsogram(word string) bool {
	m := make(map[rune]bool)

	for _, c := range word {
		if !unicode.IsLetter(c) {
			continue
		}

		c = unicode.ToLower(c)
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}
