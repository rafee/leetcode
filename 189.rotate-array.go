/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

package golang

// @lc code=start
func rotate(nums []int, k int) {
	if len(nums) != 0 {
		copy(nums, append(nums[len(nums)-k%len(nums):], nums[0:len(nums)-k%len(nums)]...))
	}
}

// @lc code=end
