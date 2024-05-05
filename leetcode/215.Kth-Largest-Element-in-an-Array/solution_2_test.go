package leetcode215

import (
	"testing"
)

func TestProblem215_2(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{5, 2, 4, 1, 3, 6, 0},
			k:        4,
			expected: 3,
		},
		{
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			nums:     []int{2, 1},
			k:        2,
			expected: 1,
		},
		{
			nums:     []int{-1, -2, -9, 1},
			k:        4,
			expected: -9,
		},
		{
			nums:     []int{0},
			k:        1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		result := findKthLargest_2(tc.nums, tc.k)
		if result != tc.expected {
			t.Error(result)
		}
	}
}
