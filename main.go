package main

import (
	"fmt"
	"sort"
)

func solution(distance int, rocks []int, n int) int {
	sort.Ints(rocks)
	left, right := 1, distance
	var mid int
	for left <= right {
		mid = (left + right) / 2
		removed := 0
		prev := 0
		for _, rock := range rocks {
			if rock-prev < mid {
				removed++
			} else {
				prev = rock
			}
		}
		if distance-prev < mid {
			removed++
		}
		if removed > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}

func main() {
	distance := 25
	rocks := []int{2, 14, 11, 21, 17}
	n := 2
	result := solution(distance, rocks, n)
	fmt.Println(result)

}
