package grains

import (
	"errors"
)

// grainValues caches results of number of grains in each square.
var grainValues [65]uint64

func Square(number int) (uint64, error) {
	if number > 64 || number <= 0 {
		return 0, errors.New("Number must be between 1 and 64")
	}

	if grainValues[number] != 0 {
		return grainValues[number], nil
	}

	if number == 1 {
		grainValues[1] = 1
		return 1, nil
	}

	v, _ := Square(number - 1)
	grainValues[number] = 2 * v
	return grainValues[number], nil
}

func Total() uint64 {
	var s uint64
	for i := 1; i <= 64; i++ {
		s1, _ := Square(i)
		s += s1
	}
	return s
}
