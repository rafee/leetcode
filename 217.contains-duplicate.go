/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

package golang

// @lc code=start
func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, num := range nums {
		if set[num] {
			return true
		}
		set[num] = true
	}
	return false
}

// @lc code=end
