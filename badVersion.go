package main

func isBadVersion(n int) bool {
	return n >= 3
}

func FindFirstBadVersion(num int) int {
	left, right := 1, num+1

	for left < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			if mid == 1 || (mid-1 > 0 && !isBadVersion(mid-1)) {
				return mid
			} else {
				right = mid
			}
		} else {
			left = mid + 1
		}
	}
	return left
}
