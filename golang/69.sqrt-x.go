/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (32.52%)
 * Likes:    1145
 * Dislikes: 1762
 * Total Accepted:    509.7K
 * Total Submissions: 1.5M
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 * the decimal part is truncated, 2 is returned.
 *
 *
 */

package golang

// @lc code=start
func mySqrt(x int) int {
	left, right := 0, x
	for left < right {
		mid := left + (right-left+1)/2
		square := mid * mid
		if square == x {
			return mid
		} else if square < x {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

// @lc code=end
