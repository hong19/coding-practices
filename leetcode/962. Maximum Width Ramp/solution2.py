from typing import List


class Solution:
    def maxWidthRamp(self, nums: List[int]) -> int:
        max_width = 0
        stack = []

        for i in range(len(nums)):
            if not stack or nums[i] < nums[stack[-1]]:
                stack.append(i)

        j = len(nums) - 1
        while j >= 0 and stack: 
            if nums[stack[-1]] <= nums[j]:
                max_width = max(j - stack.pop(), max_width)
            else:
                j -= 1 

        return max_width
