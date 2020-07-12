/*
 * @lc app=leetcode id=645 lang=golang
 *
 * [645] Set Mismatch
 */

package leetcode

import "sort"

// @lc code=start
func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	result := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i] != (i + 1) {

		}
	}
	return result
}

// @lc code=end
