#
# @lc app=leetcode id=539 lang=python3
#
# [539] Minimum Time Difference
#

# @lc code=start
from typing import List


class Solution:
    def findMinDifference(self, timePoints: List[str]) -> int:
        # timePoints = sorted(timePoints)
        timePoints.sort()
        minDiff = 1440
        for i in range(1, len(timePoints)):
            minDiff = self.min(minDiff, self.timeDiff(
                timePoints[i], timePoints[i-1]))
        return self.min(minDiff, self.timeDiff(timePoints[0], timePoints[-1]))

    def timeDiff(self, time1: str, time2: str) -> int:
        if time1 > time2:
            time1, time2 = time2, time1
        diff = (int(time2[:2])-int(time1[:2])) * \
            60+(int(time2[3:])-int(time1[3:]))
        return self.min(diff, 1440-diff)

    def min(self, val1: int, val2: int):
        if val1 > val2:
            return val2
        return val1

# @lc code=end
