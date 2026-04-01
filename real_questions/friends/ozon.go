package main

import (
	"crypto/md5"
	"fmt"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}

func RecoverPassword(h []byte) string {
	for i := 1; i <= len(h); i++ {
		password := find(i, h)
		if password != "" {
			return password
		}
	}
	return ""
}

// func TestRecoverPassword(t *testing.T) {
// 	for _, exp := range []string{
// 		"a",
// 		"12",
// 		"abc333d",
// 	} {
// 		t.Run(exp, func(t *testing.T) {
// 			act := RecoverPassword(hashPassword(exp))
// 			if act != exp {
// 				t.Error("recovered:", act, "expected:", exp)
// 			}
// 		})
// 	}
// }

func hashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}

func check(h []byte, password string) bool {
	newPass := hashPassword(password)

	for i, v := range newPass {
		if v != h[i] {
			return false
		}
	}

	return true
}

func find(length1 int, h []byte) string {
	current := make([]rune, length1)

	var recurs func(pos int) bool

	recurs = func(pos int) bool {
		if pos == length1 {
			return check(h, string(current))
		}

		for _, i := range alphabet {
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

func main() {
	for _, exp := range []string{
		"a",
		"12",
		"abc333d",
	} {
		act := RecoverPassword(hashPassword(exp))
		if act != exp {
			fmt.Println("recovered:", act, "expected:", exp)
		}
	}
}
