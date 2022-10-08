package scrabble

import (
	"unicode"
)

func createPoints() map[rune]int {
	points := map[rune]int{}

	for _, char := range "aeioulnrst" {
		points[char] = 1
	}

	for _, char := range "dg" {
		points[char] = 2
	}

	for _, char := range "bcmp" {
		points[char] = 3
	}

	for _, char := range "fhvwy" {
		points[char] = 4
	}

	for _, char := range "k" {
		points[char] = 5
	}

	for _, char := range "jx" {
		points[char] = 8
	}

	for _, char := range "qz" {
		points[char] = 10
	}

	return points
}

func Score(word string) int {
	points := createPoints()
	s := 0
	for _, c := range word {
		s += points[unicode.ToLower(c)]
	}
	return s
}
