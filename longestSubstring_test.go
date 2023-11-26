package main

import (
	"fmt"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	testCases := []struct {
		given    string
		expected int
	}{
		{"pwwkew", 3},
		{"aaaaaaaa", 1},
		{"abcabcbb", 3},
		{"abcxyzab", 6},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestMaxArea-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := LongestSubstring(testCase.given)
			if actual != testCase.expected {
				t.Errorf("Failed: expected=%d, actual=%d", testCase.expected, actual)
			}
		})
	}
}
