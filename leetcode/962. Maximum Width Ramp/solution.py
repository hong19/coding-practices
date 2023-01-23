from typing import List


class Solution:
    def maxWidthRamp(self, nums: List[int]) -> int:
        max_width = 0
        stack = []
        
        for i in range(len(nums)):
            if not stack:
                stack.append((nums[i], i))
            else:
                if nums[i] < stack[-1][0]:
                    stack.append((nums[i], i))

            upper = len(stack) - 1
            lower = 0
            target = nums[i] 
            while lower < upper:
                current = int((upper + lower)/2)
                if target >= stack[current][0]:
                    upper = current 
                else: 
                    lower = current + 1
                # print(f"target: {target}, stack: {stack}, current: {current}, upper: {upper}, lower: {lower}")
        
            if (width := i - stack[lower][1]) > max_width:
                max_width = width

        
        return max_width
