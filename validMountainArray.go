package main

func ValidMountainArray(A []int) bool {
	i := 1
	for i < len(A) && A[i] > A[i-1] {
		i++
	}
	if i == 1 || i == len(A) {
		return false
	}
	for i < len(A) && A[i] < A[i-1] {
		i++
	}

	return i == len(A)
}
