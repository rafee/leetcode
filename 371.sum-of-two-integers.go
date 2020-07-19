/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */

package golang

// @lc code=start
func getSum(a int, b int) int {
	for b != 0 {
		c := a & b
		a = a ^ b
		b = c << 1
	}
	return a
}

// @lc code=end
