package raindrops

import "strconv"

func Convert(number int) string {
	isFac := false
	s := ""
	if number%3 == 0 {
		s += "Pling"
		isFac = true
	}

	if number%5 == 0 {
		s += "Plang"
		isFac = true
	}

	if number%7 == 0 {
		s += "Plong"
		isFac = true
	}

	if !isFac {
		return strconv.Itoa(number)
	} else {
		return s
	}

}
