package isogram

import (
	"strings"
	"unicode"
)

var exists = struct{}{}

// IsIsogram returns True when a word is an isogram and False otherwise.
func IsIsogram(word string) bool {
	s := strings.ToLower(word)

	for i, c := range s {
		if unicode.IsLetter(c) && strings.ContainsRune(s[i+1:], c) {
			return false
		}
	}
	return true
}
