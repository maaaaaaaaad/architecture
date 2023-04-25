package main

import "fmt"

func binarySearch(arr []int, low, high, target int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearch(arr, mid+1, high, target)
	} else {
		return binarySearch(arr, low, mid-1, target)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 6

	index := binarySearch(arr, 0, len(arr)-1, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found in the array\n", target)
	}
}
