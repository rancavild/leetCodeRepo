package main

func SearchRangeBruteForce(input []int, target int) []int {
	result := []int{-1, -1}
	j := 1
	for i, v := range input {
		if v <= target && v == target {
			if j == 1 {
				result[0] = i
			}
			j++
		}
	}
	result[1] = j
	return result
}

func SearchRangeFirstApproach(input []int, target int) []int {
	first := -1
	for i := 0; i < len(input); i++ {
		if input[i] == target {
			first = i
			break
		}
	}
	last := -1
	for i := len(input) - 1; i > 0; i-- {
		if input[i] == target {
			last = i
			break
		}
	}

	return []int{first, last}
}

func searchFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func searchLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	first := searchFirst(nums, target)
	last := searchLast(nums, target)

	return []int{first, last}
}
