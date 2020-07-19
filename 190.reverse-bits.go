/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

package golang

// @lc code=start
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res <<= 1
		bit := num & 1
		num >>= 1
		res |= bit
	}
	return res
}

// @lc code=end
