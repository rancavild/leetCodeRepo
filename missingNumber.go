package main

func MissingNumber(nums []int) int {
	n := len(nums)
	intendedSum := n * (n + 1) / 2

	numsSum := 0
	for i := 0; i < n; i++ {
		numsSum += nums[i]
	}

	return intendedSum - numsSum
}
