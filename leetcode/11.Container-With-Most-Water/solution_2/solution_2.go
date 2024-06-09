package leetcode11

func maxArea(height []int) int {
	max := 0

	left := 0
	right := len(height) - 1

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}

	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
