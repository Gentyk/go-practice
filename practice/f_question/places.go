package main

import "fmt"

func getPlace(places []int) int {
	n := len(places)
	l_max, r_max := -1, -1
	l_ind, r_ind := 0, n-1
	i := 0
	for places[i] == 0 {
		i++
	}
	if i != 0 {
		l_max = i
	}
	l_ind = i
	i = n - 1
	for places[i] == 0 {
		i--
	}
	if i != n-1 {
		r_max = n - 1 - i
	}
	r_ind = i

	max_per, ind_max_per := 0, -1
	_per, ind_per := 0, -1
	for i := l_ind; i < r_ind; i++ {
		if places[i] == 0 {
			if _per == 0 {
				ind_per = i
			}
			_per++
		} else {
			if _per > 0 && _per > max_per {
				max_per = _per
				ind_max_per = ind_per
				_per = 0
				ind_per = -1
			}
		}
	}

	fmt.Println(l_max, max_per, r_max)

	if l_max > 0 && l_max-1 > max_per/2 {
		ind_max_per = 0
		max_per = l_max
	}
	if r_max > 0 && r_max-1 > max_per/2 {
		ind_max_per = n - 1
		max_per = r_max
	}

	if ind_max_per == 0 {
		return 0
	}
	if ind_max_per == n-1 {
		return n - 1
	}
	return ind_max_per + max_per/2

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
	checker([]int{1, 1, 0, 0, 1, 0, 0}, 6)
	checker([]int{1, 1, 0, 0, 0, 1, 0}, 3)
	checker([]int{1, 1, 0, 0, 1}, 3)
}
