from typing import List
import math


class Solution:
    def maxFrequency(self, nums: List[int], k: int) -> int:
        nums = sorted(nums)
        
        prefix = [0]
        
        for i in reversed(nums):
            prefix.append(prefix[-1] + nums[-1] - i)

        upper = len(prefix) - 1
        lower = 0
    
        # print(f"prefix: {prefix}")
        for value in reversed(nums):
            

        while upper > lower: 
            current = math.ceil((upper + lower)/2)
            # print(f"u: {upper}, l: {lower}, current: {current}")
            if prefix[current] > k : 
                upper = current - 1
            else: 
                lower = current

        print(lower) 

        return lower