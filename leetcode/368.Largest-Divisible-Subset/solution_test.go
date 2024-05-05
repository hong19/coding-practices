package leetcode368

import (
	"fmt"
	"testing"
)

type testCase struct {
	nums     []int
	expected []int
}

func TestLargestDivisibleSubset(t *testing.T) {
	testCases := []testCase{
		{
			[]int{1, 2, 3},
			[]int{2, 1},
		},
		{
			[]int{1, 2, 4, 8},
			[]int{8, 4, 2, 1},
		},
		{
			[]int{3, 1, 2, 4, 8},
			[]int{8, 4, 2, 1},
		},
		{
			[]int{11},
			[]int{11},
		},
		{
			[]int{4, 8, 10, 240},
			[]int{240, 8, 4},
		},
		{
			[]int{3, 4, 16, 8},
			[]int{16, 8, 4},
		},
	}

	for _, tc := range testCases {
		results := largestDivisibleSubset(tc.nums)

		if len(tc.expected) != len(results) {
			t.Errorf("result len: %d, expected len: %d\n", len(results), len(tc.expected))
		}

		for i, value := range tc.expected {
			if results[i] != value {
				fmt.Printf("result: %d, expected: %d\n", results[i], value)
				t.Error()
			}
		}

	}

}
