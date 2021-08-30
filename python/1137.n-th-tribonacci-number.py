#
# @lc app=leetcode id=1137 lang=python3
#
# [1137] N-th Tribonacci Number
#

# @lc code=start
class Solution:
    def tribonacci(self, n: int) -> int:
        if n < 2:
            return n
        elif n == 2:
            return 1
        trib1, trib2, trib3 = 0, 1, 1
        i = 3
        while i <= n:
            trib3, trib2, trib1 = trib3+trib2+trib1, trib3, trib2
            i += 1
        return trib3

# @lc code=end
