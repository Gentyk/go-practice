package main

import "fmt"

/**
 * Места в кинотеатре расположены в один ряд. Только что пришедший зритель выбирает место,
 * чтобы сидеть максимально далеко от остальных зрителей в ряду. То есть расстояние от того места, куда сядет зритель
 * до ближайшего к нему зрителя должно быть максимально.
 * Гарантируется, что в ряду всегда есть свободные места и уже сидит хотя бы один зритель.
 * Напиши функцию, которая по заданному ряду мест (массиву из нулей и единиц) вернёт расстояние
 * от выбранного места до ближайшего зрителя.
 * Input: [1, 0, 0, 0, 1]
 * Output: 2
 * <p>
 * Input: [1, 0, 1, 0, 0, 1, 0, 0, 0, 1]
 * Output: 2
 * <p>
 * Input: [1, 0, 1, 0]
 * Output: 1
  * <p>
 * Input: [0, 0, 0, 1, 0, 1, 0]
 * Output: 3
  * <p>
 * Input: [1, 0, 1, 0, 0]
 * Output: 2
 */

func findPlace(pl []int) int {
    result := 0
    n := len(pl)
    left := 0
    right := n-1
    
    // left
    size := 0
    for i:=0; i<n; i++ {
        if pl[i] == 0 {
            size = size + 1
        } else {
            left = i
            break
        }
    }
    result = size
    
    // right
    size = 0
    for i:=n-1; i>0; i-- {
        if pl[i] == 0 {
            size = size + 1
        } else {
            right = i
            break
        }
    }
    if result < size {
        result = size
    }
    size = 0
    
    // medium
    for i := left; i <= right; i++ {
        if pl[i] == 1 {
            if size != 0 {
                if result < (size / 2 + size % 2) {
                    result = (size / 2 + size % 2)
                }
                size = 0
            }
        } else {
            size = size + 1
        }
    }
    
    return result 
}


func assert(seats []int, expected int) {
	maxDist := findPlace(seats)
	if maxDist != expected {
		fmt.Printf("FAIL %v - %d, expected %d \n", seats, maxDist, expected)
	} else {
		fmt.Printf("OK %v - %d \n", seats, maxDist)
	}
}

func main() {
	assert([]int{1, 0, 0, 0, 1}, 2)
	assert([]int{1, 0, 1, 0, 0, 1, 0, 0, 0, 1}, 2)
	assert([]int{1, 0, 1, 0}, 1)
	assert([]int{0, 0, 0, 1, 0}, 3)
	assert([]int{1, 0, 0, 0, 0}, 4)
}
