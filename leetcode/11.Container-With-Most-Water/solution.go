package leetcode11

func maxArea(height []int) int {
	max := 0

	for left := 0; left < len(height); left++ {
		for right := len(height) - 1; right > left; right-- {
			area := min(height[left], height[right]) * (right - left)
			// fmt.Printf("left: %d %d, right: %d %d, area: %d\n", left, height[left], right, height[right], area)
			if area > max {
				max = area
			}
			if height[left]*area <= max {
				break
			}
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
