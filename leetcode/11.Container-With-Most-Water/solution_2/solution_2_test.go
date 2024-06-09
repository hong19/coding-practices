package leetcode11

import (
	"testing"
)

type testCase struct {
	height   []int
	expected int
}

func TestCoinChange(t *testing.T) {
	testCases := []testCase{
		{
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			[]int{1, 1},
			1,
		},
	}

	for _, tc := range testCases {
		result := maxArea(tc.height)

		if tc.expected != result {
			t.Errorf("result: %d, expected: %d\n", result, tc.expected)
		}
	}
}
