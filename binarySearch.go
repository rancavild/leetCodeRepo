package main

import (
	"sort"
)

func BinarySearch(arr []int, target int) int {
	sort.Ints(arr)
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
