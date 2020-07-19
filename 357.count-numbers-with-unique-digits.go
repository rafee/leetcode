/*
 * @lc app=leetcode id=357 lang=golang
 *
 * [357] Count Numbers with Unique Digits
 */

package leetcode

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 10
	}
	sum := 10
	count := 9
	for i := 0; i < n-1; i++ {
		count *= (9 - i)
		sum += count
	}
	return sum
}

// @lc code=end
