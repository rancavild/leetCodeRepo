package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		given    []int
		expected []int
	}{
		{[]int{0, 1, 0, 2, 3}, []int{1, 2, 3, 0, 0}},
		{[]int{0, 0, 12, 2, 3}, []int{12, 2, 3, 0, 0}},
		{[]int{0}, []int{0}},
	}

	for i, testCase := range testCases {
		testName := fmt.Sprintf("TestMaxArea-%d", i)
		t.Run(testName, func(t *testing.T) {
			MoveZeroes(testCase.given)
			if !reflect.DeepEqual(testCase.given, testCase.expected) {
				t.Errorf("Failed: expected=%v, actual=%v", testCase.expected, testCase.given)
			}
		})
	}
}
