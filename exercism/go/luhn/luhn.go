package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	s := 0
	for i := 0; i < len(id); i++ {
		d, ok := strconv.Atoi(string(id[len(id)-1-i]))
		if ok != nil {
			return false
		}

		if i%2 == 1 {
			if d == 9 {
				s += d
			} else {
				s += (2 * d) % 9
			}
		} else {
			s += d
		}
	}

	if len(id) > 1 && s%10 == 0 {
		return true
	}
	return false
}
