package leetcode80

import "fmt"

func removeDuplicates(nums []int) int {
	cur_count := 1
	cur_idx := 1
	output_idx := 1

	if len(nums) <= 2 {
		return len(nums)
	}

	for cur_idx = 1; cur_idx < len(nums); cur_idx++ {
		if nums[cur_idx] == nums[cur_idx-1] {
			cur_count++
		} else {
			cur_count = 1
		}

		if cur_count <= 2 {
			nums[output_idx] = nums[cur_idx]
			// fmt.Printf("cur_idx: %d, output num: %d output idx:%d \n", cur_idx, nums[output_idx], output_idx)
			output_idx++
		}
	}

	return output_idx
}
