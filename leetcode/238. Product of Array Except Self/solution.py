from typing import List

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        prefix = []
        suffix = []

        ans = []

        for i in range(len(nums)):
            if prefix:
                prefix.append(prefix[-1] * nums[i])
            else: 
                prefix.append(nums[i])

        for i in range(len(nums))[::-1]:
            if suffix:
                suffix.append(suffix[-1] * nums[i])
            else: 
                suffix.append(nums[i])

        suffix = suffix[::-1]
    
        # print(f"prefix: {prefix}, suffix: {suffix}")
        ans.append(suffix[1])
        for i in range(1, len(nums) - 1):
            ans.append(prefix[i - 1] * suffix[i + 1])
        ans.append(prefix[-2])

        return ans 