package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		given    []int
		target   int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 7}, 1, 0},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 5, 4},
		{[]int{1, 1}, 1, 0},
		{[]int{2, 1, 0}, 3, -1},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestMaxArea-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := BinarySearch(testCase.given, testCase.target)
			if actual != testCase.expected {
				t.Errorf("Failed: expected=%d, actual=%d", testCase.expected, actual)
			}
		})
	}
}
