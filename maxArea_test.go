package main

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		given    []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 7}, 28},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestMaxArea-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := MaxArea(testCase.given)
			if actual != testCase.expected {
				t.Errorf("Failed: expected=%d, actual=%d", testCase.expected, actual)
			}
		})
	}
}
