package main

import (
	"fmt"
	"testing"
)

func TestNumSquares(t *testing.T) {
	testCases := []struct {
		given    int
		expected int
	}{
		{12, 3},
		{13, 2},
		{100, 1},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestMaxArea-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := NumSquares(testCase.given)
			if actual != testCase.expected {
				t.Errorf("Failed: expected=%d, actual=%d", testCase.expected, actual)
			}
		})
	}
}
