/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

package golang

// @lc code=start
func hammingDistance(x int, y int) int {
	count := 0
	for x != 0 || y != 0 {
		count += (x & 1) ^ (y & 1)
		x >>= 1
		y >>= 1
	}
	return count
}

// @lc code=end
