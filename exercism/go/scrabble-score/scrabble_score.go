package scrabble

import (
	"unicode"
)

func createPoints() map[string]int {
	points := map[string]int{}

	for _, char := range "aeioulnrst" {
		points[string(char)] = 1
	}

	for _, char := range "dg" {
		points[string(char)] = 2
	}

	for _, char := range "bcmp" {
		points[string(char)] = 3
	}

	for _, char := range "fhvwy" {
		points[string(char)] = 4
	}

	for _, char := range "k" {
		points[string(char)] = 5
	}

	for _, char := range "jx" {
		points[string(char)] = 8
	}

	for _, char := range "qz" {
		points[string(char)] = 10
	}

	return points
}

func Score(word string) int {
	points := createPoints()
	s := 0
	for _, c := range word {
		s += points[string(unicode.ToLower(c))]
	}
	return s
}
