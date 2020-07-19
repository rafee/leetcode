/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 *
 * https://leetcode.com/problems/maximal-square/description/
 *
 * algorithms
 * Medium (34.80%)
 * Likes:    2610
 * Dislikes: 65
 * Total Accepted:    228.1K
 * Total Submissions: 621.9K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * Given a 2D binary matrix filled with 0's and 1's, find the largest square
 * containing only 1's and return its area.
 *
 * Example:
 *
 *
 * Input:
 *
 * 1 0 1 0 0
 * 1 0 1 1 1
 * 1 1 1 1 1
 * 1 0 0 1 0
 *
 * Output: 4
 *
 */

package golang

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	dp := make([]int, cols+1)
	maxsqlen, prev := 0, 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			temp := dp[j]
			if matrix[i-1][j-1] == '1' {
				dp[j] = min(min(dp[j-1], prev), dp[j]) + 1
				maxsqlen = max(maxsqlen, dp[j])
			} else {
				dp[j] = 0
			}
			prev = temp
		}
	}
	return maxsqlen * maxsqlen
}

// @lc code=end
