package isogram

import (
	"strings"
	"unicode"
)

var exists = struct{}{}

// IsIsogram returns True when a word is an isogram and False otherwise.
func IsIsogram(word string) bool {
	// mimic a set with a 0 byte type as value.
	s := make(map[rune]struct{})

	for _, c := range word {
		c = unicode.ToLower(c)
		if strings.ContainsRune(" -", c) {
			continue
		}

		if _, ok := s[c]; ok {
			return false
		}
		s[c] = exists
	}
	return true
}
