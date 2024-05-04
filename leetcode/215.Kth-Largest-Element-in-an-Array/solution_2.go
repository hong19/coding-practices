package leetcode215

/*
Quick select
*/

func findKthLargest_2(nums []int, k int) int {
	return quickSelect(nums, k, 0, len(nums))
}

func quickSelect(nums []int, k int, start int, end int) int {
	pivot := nums[start]
	nextPivotIdx := start + 1

	if (start - end) == 1 {
		return pivot
	}
	for i := start; i < end; i++ {
		value := nums[i]
		if value > pivot {
			// swap the num to left partition
			nums[i], nums[nextPivotIdx] = nums[nextPivotIdx], nums[i]
			nextPivotIdx++
		}
	}
	// swap pivot with last swapped value
	currentPivotIdx := nextPivotIdx - 1
	nums[start] = nums[currentPivotIdx]
	nums[currentPivotIdx] = pivot

	// fmt.Printf("array: %v, start: %d, end: %d, currentPivotIdx: %d, k: %d, pivot: %d\n", nums, start, end, currentPivotIdx, k, pivot)
	if currentPivotIdx+1 == k {
		return pivot
	} else if currentPivotIdx+1 > k {
		return quickSelect(nums, k, start, currentPivotIdx)
	} else {
		return quickSelect(nums, k, currentPivotIdx+1, end)
	}
}
