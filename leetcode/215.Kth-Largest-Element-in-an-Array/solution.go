package leetcode215

/*
Binary search for the answer
*/

func findKthLargest(nums []int, k int) int {
	upper, lower := getMaxMin(nums)

	for upper > lower {
		target := (upper + lower)
		if target > 0 {
			target = target / 2
		} else {
			// handle negative int
			target = target/2 - 1
		}
		if isKNumLargerThanTarget(nums, k, target) {
			lower = target + 1
		} else {
			upper = target
		}
	}

	return upper
}

func isKNumLargerThanTarget(nums []int, k int, target int) bool {
	count := 0

	for _, value := range nums {
		if value > target {
			count++
			if count >= k {
				return true
			}
		}
	}

	return false
}

func getMaxMin(nums []int) (int, int) {
	max := nums[0]
	min := nums[0]

	for _, value := range nums {
		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}
	return max, min
}
