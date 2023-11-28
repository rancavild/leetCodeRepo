package main

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	testCases := []struct {
		given    []int
		expected int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestMaxArea-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := MissingNumber(testCase.given)
			if actual != testCase.expected {
				t.Errorf("Failed: expected=%d, actual=%d", testCase.expected, actual)
			}
		})
	}
}
