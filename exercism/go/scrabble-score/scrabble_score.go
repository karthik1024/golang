package scrabble

import (
	"errors"
	"fmt"
	"strings"
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

//var POINTS = createPoints()

func charScore(char rune) (int, error) {
	if strings.ContainsRune("aeioulnrst", char) {
		return 1, nil
	} else if strings.ContainsRune("dg", char) {
		return 2, nil
	} else if strings.ContainsRune("bcmp", char) {
		return 3, nil
	} else if strings.ContainsRune("fhvwy", char) {
		return 4, nil
	} else if strings.ContainsRune("k", char) {
		return 5, nil
	} else if strings.ContainsRune("jx", char) {
		return 8, nil
	} else if strings.ContainsRune("qz", char) {
		return 10, nil
	} else {
		return 0, errors.New(fmt.Sprintf("Unknown rune %c", char))
	}
}

// Score returns the scrabble score for a word.
func Score(word string) int {
	s := 0

	for _, c := range word {
		cs, err := charScore(unicode.ToLower(c))
		if err == nil {
			s += cs
		}
	}
	return s
}
