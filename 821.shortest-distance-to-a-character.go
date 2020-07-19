/*
 * @lc app=leetcode id=821 lang=golang
 *
 * [821] Shortest Distance to a Character
 */

package golang

// @lc code=start
func shortestToChar(S string, C byte) []int {
	dist := make([]int, len(S))
	pos := -len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			pos = i
		}
		dist[i] = i - pos
	}

	pos = 2 * len(S)
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == C {
			pos = i
		}
		dist[i] = min(dist[i], pos-i)
	}

	return dist
}

// @lc code=end
