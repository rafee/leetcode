/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

package golang

// @lc code=start
func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}

// @lc code=end
