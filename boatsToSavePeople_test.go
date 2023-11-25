package main

import (
	"fmt"
	"testing"
)

func TestNumRescueBoats(t *testing.T) {
	testCases := []struct {
		people   []int
		limit    int
		expected int
	}{
		{[]int{1, 2}, 3, 1},
		{[]int{3, 2, 2, 1}, 3, 3},
		{[]int{3, 5, 3, 4}, 5, 4},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestMaxArea-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := NumRescueBoats(testCase.people, testCase.limit)
			if actual != testCase.expected {
				t.Errorf("Failed: expected=%d, actual=%d", testCase.expected, actual)
			}
		})
	}
}
