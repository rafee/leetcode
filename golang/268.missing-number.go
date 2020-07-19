/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

package golang

// @lc code=start
func missingNumber(nums []int) int {
	sum := len(nums) * (len(nums) + 1) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}

// @lc code=end
