/*
 * @lc app=leetcode id=886 lang=golang
 *
 * [886] Possible Bipartition
 */

package leetcode

// @lc code=start
func possibleBipartition(N int, dislikes [][]int) bool {
	graph := make([][]bool, N)
	for i := 0; i < N; i++ {
		graph[i] = make([]bool, N)
	}

	for _, dislike := range dislikes {
		graph[dislike[0]-1][dislike[1]-1] = true
		graph[dislike[1]-1][dislike[0]-1] = true
	}
	colors := make(map[int]bool)

	for i := 0; i < N; i++ {
		_, found := colors[i]
		if !found && !visitNext(graph, colors, i, true) {
			return false
		}
	}
	return true
}

func visitNext(graph [][]bool, colors map[int]bool, src int, color bool) bool {
	colors[src] = color
	for i := 0; i < len(graph); i++ {
		if graph[src][i] == true {
			_, found := colors[i]
			if found && colors[i] == color {
				return false
			}
			if !found && !visitNext(graph, colors, i, !color) {
				return false
			}
		}
	}
	return true
}

// @lc code=end
