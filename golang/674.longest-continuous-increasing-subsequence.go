/*
 * @lc app=leetcode id=674 lang=golang
 *
 * [674] Longest Continuous Increasing Subsequence
 */

package golang

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mx, cur := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cur++
		} else {
			if cur > mx {
				mx = cur
			}
			cur = 1
		}
	}
	if cur > mx {
		mx = cur
	}
	return mx
}

// @lc code=end
