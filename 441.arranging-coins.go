/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 */

package leetcode

// @lc code=start
func arrangeCoins(n int) int {
	left, right := 0, n
	for left <= right {
		k := left + (right-left)/2
		cur := k * (k + 1) / 2
		if cur == n {
			return k
		} else if cur < n {
			left = k + 1
		} else {
			right = k - 1
		}
	}
	return right
}

// @lc code=end
