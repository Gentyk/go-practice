package main

import "fmt"

func getPlace(places []int) int {
	n := len(places)
	left := 0
	right := n - 1

	for i, pl := range places {
		if pl == 1 {
			if i-1 > 0 {
				left = i
			}
			break
		}
	}

	for i := n - 1; i > 0; i-- {
		if places[i] == 1 {
			if i+1 < right {
				right = i
			}
			break
		}
	}
	ind := 0
	max := left - 1
	if n-1 != right && max < n-1-right-1 {
		ind = n - 1
		max = n - 1 - right - 1
	}

	left, right = 0, 0
	dist2 := 0
	for i := left; i < right; i++ {
		if places[i] == 1 {
			if left != 0 {
				dist2 = (right-left)/2 - 1 + (right-left)%2
				if dist2 > max {
					max = dist2
					ind = left + dist2
				}
				left, right = 0, 0
			}
		} else {
			if left == 0 {
				left, right = i, i
			} else {
				right = i
			}
		}
	}

	return ind

}

func checker(d []int, resp int) {
	if getPlace(d) != resp {
		fmt.Println("Ошибка в тесте ", d, ": ", getPlace(d), "!=", resp)
	} else {
		fmt.Println("ok")
	}

}

func main() {
	checker([]int{1, 1, 0, 0}, 3)
	checker([]int{0, 1, 1, 0, 0}, 4)
	checker([]int{0, 0, 1, 1, 0}, 0)
	checker([]int{1, 1, 0, 0, 1, 0, 0}, 6)
	checker([]int{1, 1, 0, 0, 0, 1, 0}, 3)
	checker([]int{1, 1, 0, 0, 1}, 3)
}
