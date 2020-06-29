/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

package leetcode

import "math"

// @lc code=start
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt64
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

// @lc code=end
