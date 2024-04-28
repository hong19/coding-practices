package leetcode80

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}      // Input array
	expectedNums := []int{1, 1, 2, 2, 3} // The expected answer with correct length

	k := removeDuplicates(nums) // Calls your implementation
	for i := 0; i < k; i++ {
		if nums[i] != expectedNums[i] {
			t.Error()
			break
		}
	}
}
