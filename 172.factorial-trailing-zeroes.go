/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

package leetcode

// @lc code=start
func trailingZeroes(n int) int {
	sum := 0
	for n > 0 {
		n /= 5
		sum += n
	}
	return sum
}

// @lc code=end
