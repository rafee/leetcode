/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

package leetcode

// @lc code=start
func isPerfectSquare(num int) bool {
	left, right := 0, num
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		mid2 := mid * mid
		if mid2 == num {
			return true
		} else if mid2 > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// @lc code=end
