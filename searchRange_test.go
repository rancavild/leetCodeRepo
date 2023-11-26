package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct {
		given    []int
		target   int
		expected []int
	}{
		{[]int{1, 2, 3, 3, 3, 4}, 3, []int{2, 4}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{2, 2}},
		{[]int{1, 2, 3, 4, 5}, 10, []int{-1, -1}},
		{[]int{1, 2, 11, 11, 11, 40}, 11, []int{2, 4}},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestMaxArea-%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := SearchRange(testCase.given, testCase.target)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Failed: expected=%v, actual=%v", testCase.expected, actual)
			}
		})
	}
}
