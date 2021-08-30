#
# @lc app=leetcode id=977 lang=python3
#
# [977] Squares of a Sorted Array
#

# @lc code=start
from typing import List


class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        nums = sorted(nums, key=lambda x: abs(x))
        result = [None]*len(nums)
        for i in range(len(nums)):
            result[i] = nums[i]**2
        return result

# @lc code=end
