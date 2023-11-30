package main

import (
	"fmt"
	"testing"
)

func TestNearestPalindrome(t *testing.T) {
	testCases := []struct {
		given    string
		expected string
	}{
		{"123", "121"},
		{"1", "0"},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestMaxArea-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := NearestPalindrome(testCase.given)
			if actual != testCase.expected {
				t.Errorf("Failed: expected=%s, actual=%s", testCase.expected, actual)
			}
		})
	}
}
