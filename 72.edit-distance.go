/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

package leetcode

// @lc code=start
func minDistance(word1 string, word2 string) int {
	W1, W2 := len(word1), len(word2)
	dp := make([][]int, W1+1)
	for i := 0; i <= W1; i++ {
		dp[i] = make([]int, W2+1)
		dp[i][0] = i
	}
	for i := 0; i <= W2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= W1; i++ {
		for j := 1; j <= W2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[W1][W2]
}

// @lc code=end
