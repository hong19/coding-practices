package leetcode215

import (
	"testing"
)

func TestProblem215(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
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
		result := findKthLargest(tc.nums, tc.k)
		if result != tc.expected {
			t.Error(result)
		}
	}
}

func TestIsGreaterThanKthLargest(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		target   int
		expected bool
	}{
		{
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			target:   3,
			expected: true,
		},
		{
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			target:   5,
			expected: false,
		},
	}

	for _, tc := range testCases {
		result := isKNumLargerThanTarget(tc.nums, tc.k, tc.target)
		if result != tc.expected {
			t.Error(result)
		}
	}
}
