package main

func MoveZeroes(nums []int) {
	j := 0
	for _, num := range nums {
		if num != 0 {
			nums[j] = num
			j++
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}
