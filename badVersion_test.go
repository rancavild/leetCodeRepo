package main

import (
	"fmt"
	"testing"
)

func TestFindFirstBadVersion(t *testing.T) {
	testCases := []struct {
		given    int
		expected int
	}{
		{4, 3},
		{5, 3},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestMaxArea-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := FindFirstBadVersion(testCase.given)
			if actual != testCase.expected {
				t.Errorf("Failed: expected=%d, actual=%d", testCase.expected, actual)
			}
		})
	}
}
