#
# @lc app=leetcode id=278 lang=python3
#
# [278] First Bad Version
#

# @lc code=start
# The isBadVersion API is already defined for you.
# @param version, an integer
# @return an integer
# def isBadVersion(version):

class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        good, bad = 1, n
        while good < bad:
            mid = good+(bad-good)//2
            if isBadVersion(mid):
                bad = mid
            else:
                good = mid+1
        return bad

# @lc code=end
