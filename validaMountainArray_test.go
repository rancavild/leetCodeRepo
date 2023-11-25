package main

import (
	"fmt"
	"testing"
)

func TestValidMountainArray(t *testing.T) {
	testCases := []struct {
		given    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 2, 1, 0}, true},
		{[]int{1, 2, 3}, false},
		{[]int{1, 1, 2, 3, 1, 3}, false},
		{[]int{10, 20, 10, 5}, true},
		{[]int{30, 20, 10, 5}, false},
		{[]int{}, false},
		{[]int{1}, false},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestValidMountainArray-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := ValidMountainArray(testCase.given)
			if actual != testCase.expected {
				t.Errorf("Failed: expected=%v, actual=%v", testCase.expected, actual)
			}
		})
	}
}
