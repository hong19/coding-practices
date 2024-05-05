package leetcode368

import (
	// "fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	maxGroupCount := make([]int, len(nums))
	prevIdx := make([]int, len(nums))

	for i := range maxGroupCount {
		maxGroupCount[i] = 1
	}

	for current := 1; current < len(nums); current++ {
		for i := current - 1; i >= 0; i-- {
			if nums[current]%nums[i] == 0 && (maxGroupCount[i]+1) > maxGroupCount[current] {
				maxGroupCount[current] = maxGroupCount[i] + 1
				prevIdx[current] = i
			}
		}
	}
	maxGroupIdx := findMaxIdx(maxGroupCount)

	ans := []int{}
	maxSize := maxGroupCount[maxGroupIdx]
	for i := maxGroupIdx; i >= 0 && maxSize > 0; i = prevIdx[i] {
		ans = append(ans, nums[i])
		maxSize--
	}

	// fmt.Printf("nums: %v, maxGroupCount: %v, prevIdx: %v\n", nums, maxGroupCount, prevIdx)
	// fmt.Printf("ans: %v\n", ans)
	return ans

}

func findMaxIdx(nums []int) int {
	maxIdx := 0

	for i, value := range nums {
		if value > nums[maxIdx] {
			maxIdx = i
		}
	}

	return maxIdx
}
