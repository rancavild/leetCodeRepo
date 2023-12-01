package main

import (
	"fmt"
	"testing"
)

func TestPrimeNumber(t *testing.T) {
	testCases := []struct {
		given    int
		expected int
	}{
		{10, 4},
		{34, 11},
		{0, 0},
		{1, 0},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestMaxArea-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := PrimeNumber(testCase.given)
			if actual != testCase.expected {
				t.Errorf("Failed: expected=%d, actual=%d", testCase.expected, actual)
			}
		})
	}
}
