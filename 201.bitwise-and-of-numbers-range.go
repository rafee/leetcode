/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 *
 * https://leetcode.com/problems/bitwise-and-of-numbers-range/description/
 *
 * algorithms
 * Medium (36.89%)
 * Likes:    695
 * Dislikes: 95
 * Total Accepted:    110.3K
 * Total Submissions: 293.8K
 * Testcase Example:  '5\n7'
 *
 * Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND
 * of all numbers in this range, inclusive.
 *
 * Example 1:
 *
 *
 * Input: [5,7]
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: [0,1]
 * Output: 0
 */

package leetcode

// @lc code=start
func rangeBitwiseAnd(m int, n int) int {
	for n > m {
		n &= (n - 1)
	}
	return n
}

// @lc code=end
