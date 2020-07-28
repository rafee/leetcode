/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

package golang

// @lc code=start
func addDigits(num int) int {
	sum := 0
	for ; num > 0; num /= 10 {
		sum += (num % 10)
	}
	if sum >= 10 {
		return addDigits(sum)
	}
	return sum
}

// @lc code=end
