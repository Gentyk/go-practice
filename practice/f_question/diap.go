// диапазоны go

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func diap(arr []int) string {
	if len(arr) == 0 {
		return ""
	}
	result := ""
	sort.Ints(arr)
	start := arr[0]
	end := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-end <= 1 {
			end = arr[i]
		} else {
			if result != "" {
				result += ","
			}
			if start == end {
				result += strconv.Itoa(start)
			} else {
				result += strconv.Itoa(start) + "-" + strconv.Itoa(end)
			}
			start, end = arr[i], arr[i]
		}
	}
	if result != "" {
		result += ","
	}
	if start == end {
		result += strconv.Itoa(start)
	} else {
		result += strconv.Itoa(start) + "-" + strconv.Itoa(end)
	}
	return result
}

func main() {
	fmt.Println(diap([]int{2, 2, 44, 43, 5}))
}
