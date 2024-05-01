package leetcode80

import (
	"testing"
)

type params struct {
	nums []int
}
type expected struct {
	nums []int
}
type testCase struct {
	params
	expected
}

func Test_removeDuplicates(t *testing.T) {
	// nums := []int{1, 1, 1, 2, 2, 3}      // Input array
	// expectedNums := []int{1, 1, 2, 2, 3} // The expected answer with correct length

	testCases := []testCase{
		{
			params{
				[]int{1, 1, 1, 2, 2, 3},
			},
			expected{
				[]int{1, 1, 2, 2, 3},
			},
		},
		{
			params{
				[]int{1},
			},
			expected{
				[]int{1},
			},
		},
		{
			params{
				[]int{},
			},
			expected{
				[]int{},
			},
		},
	}

	for _, tc := range testCases {
		k := removeDuplicates(tc.params.nums) // Calls your implementation
		for i := 0; i < k; i++ {
			if tc.params.nums[i] != tc.expected.nums[i] {
				t.Error()
				break
			}
		}
	}
}
