#
# @lc app=leetcode id=18 lang=python3
#
# [18] 4Sum
#

# @lc code=start
from typing import List


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        result=[]
        # hash=set(nums)
        sorted(nums)
        for i, num1 in enumerate(nums):
            for j,num2 in enumerate(nums[i+1:]):
                for k, num3 in enumerate(nums[i+j+1:]):
                    for num4 in nums[i+j+k+1:]:
                        if num1+num2+num3+num4 ==target:
                            result.append([num1, num2,num3, num4])
        
        return result
        
# @lc code=end

