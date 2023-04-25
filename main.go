package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 6

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found in the array\n", target)
	}
}
