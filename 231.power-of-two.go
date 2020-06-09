/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */

package leetcode

// @lc code=start
func isPowerOfTwo(n int) bool {
	return (n&(n-1) == 0) && n != 0
}

// @lc code=end
