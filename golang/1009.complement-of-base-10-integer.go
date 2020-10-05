/*
 * @lc app=leetcode id=1009 lang=golang
 *
 * [1009] Complement of Base 10 Integer
 */

package golang

// @lc code=start
func bitwiseComplement(N int) int {
	sum := 0
	for i := N; i > 0; i >>= 1 {
		sum <<= 1
		sum |= 1
	}
	return sum - N
}

// @lc code=end
