/*
 * @lc app=leetcode id=997 lang=golang
 *
 * [997] Find the Town Judge
 */

package golang

// @lc code=start
func findJudge(N int, trust [][]int) int {
	degree := make([]int, N+1)
	for _, pair := range trust {
		degree[pair[0]]--
		degree[pair[1]]++
	}

	for i := 1; i <= N; i++ {
		if degree[i] == (N - 1) {
			return i
		}
	}

	return -1
}

// @lc code=end
