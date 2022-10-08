package scrabble

import (
	"unicode"
)

// createPoints returns a map between letters and scores.
func createPoints() map[rune]int {
	groups := map[string]int{
		"aeioulnrst": 1,
		"dg":         2,
		"bcmp":       3,
		"fhvwy":      4,
		"k":          5,
		"jx":         8,
		"qz":         10,
	}

	points := map[rune]int{}

	for gp, p := range groups {
		for _, c := range gp {
			points[c] = p
		}
	}

	return points
}

// Score returns the scrabble score for a word.
func Score(word string) int {
	points := createPoints()
	s := 0
	for _, c := range word {
		s += points[unicode.ToLower(c)]
	}
	return s
}
