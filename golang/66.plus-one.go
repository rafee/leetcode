/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

package golang

// @lc code=start
func plusOne(digits []int) []int {
	digitlen := len(digits)
	result := make([]int, digitlen+1)
	carry := 1
	for i := digitlen; i > 0; i-- {
		result[i] = digits[i-1] + carry
		carry = result[i] / 10
		result[i] %= 10
	}
	if carry == 0 {
		return result[1:]
	}
	result[0] = carry
	return result
}

// @lc code=end
