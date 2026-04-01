package friends

import (
	"fmt"
)

var alf = []rune{'a', 'b', 'c'}

func find(length1 int, target string) string {
	current := make([]rune, length1)

	var recurs func(pos int) bool

	recurs = func(pos int) bool {
		if pos == length1 {
			if string(current) == target {
				return true
			}
			return false
		}

		for _, i := range alf {
			current[pos] = i
			if recurs(pos + 1) {
				return true
			}
		}

		return false
	}

	if recurs(0) {
		return string(current)
	}

	return ""
}
