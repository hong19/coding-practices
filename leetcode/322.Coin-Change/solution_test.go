package leetcode322

import (
	"testing"
)

type testCase struct {
	coins    []int
	amount   int
	expected int
}

func TestCoinChange(t *testing.T) {
	testCases := []testCase{
		{
			[]int{1, 2, 5},
			11,
			3,
		},
		{
			[]int{2},
			3,
			-1,
		},
		{
			[]int{1},
			0,
			0,
		},
	}

	for _, tc := range testCases {
		result := coinChange(tc.coins, tc.amount)

		if tc.expected != result {
			t.Errorf("result: %d, expected: %d\n", result, tc.expected)
		}
	}
}
