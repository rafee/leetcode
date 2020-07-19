/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

package golang

import "math"

// @lc code=start
func lengthOfLIS(nums []int) int {
	mx := getLISWithMin(math.MinInt64, nums)
	return mx
}

func getLISWithMin(num int, nums []int) int {
	mx := 0
	for i, cur := 0, 0; i < len(nums)-mx; i++ {
		if nums[i] > num {
			cur = getLISWithMin(nums[i], nums[i+1:]) + 1
			if cur > mx {
				mx = cur
			}
		}
	}
	return mx
}

// @lc code=end
