/*
 * @lc app=leetcode id=611 lang=golang
 *
 * [611] Valid Triangle Number
 */

package golang

import "sort"

// @lc code=start
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := len(nums) - 1; i >= 2; i-- {
		l, r := 0, i-1
		for l < r {
			if nums[l]+nums[r] > nums[i] {
				count += (r - l)
				r--
			} else {
				l++
			}
		}
	}
	return count
}

// @lc code=end
