#
# @lc app=leetcode id=509 lang=python3
#
# [509] Fibonacci Number
#

# @lc code=start
class Solution:
    def fib(self, n: int) -> int:
        if n <= 1:
            return n
        fib1, fib2 = 0, 1
        i = 2
        while i <= n:
            fib2, fib1 = fib2+fib1, fib2
            i += 1
        return fib2

# @lc code=end
