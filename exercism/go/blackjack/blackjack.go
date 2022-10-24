package blackjack

import (
	"strings"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	c := strings.ToLower(card)
	if c == "ace" {
		return 11
	} else if c == "two" {
		return 2
	} else if c == "three" {
		return 3
	} else if c == "four" {
		return 4
	} else if c == "five" {
		return 5
	} else if c == "six" {
		return 6
	} else if c == "seven" {
		return 7
	} else if c == "eight" {
		return 8
	} else if c == "nine" {
		return 9
	} else if c == "ten" {
		return 10
	} else if c == "jack" {
		return 10
	} else if c == "queen" {
		return 10
	} else if c == "king" {
		return 10
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1 := strings.ToLower(card1)
	c2 := strings.ToLower(card2)
	dc := strings.ToLower(dealerCard)
	s := ParseCard(c1) + ParseCard(c2)
	if c1 == "ace" && c2 == "ace" {
		return "P"
	}
	if s == 21 {
		if ParseCard(dc) < 10 {
			return "W"
		} else {
			return "S"
		}
	}
	if s >= 17 && s <= 20 {
		return "S"
	}
	if s >= 12 && s <= 16 {
		if ParseCard(dc) >= 7 {
			return "H"
		} else {
			return "S"
		}
	}
	return "H"
}
