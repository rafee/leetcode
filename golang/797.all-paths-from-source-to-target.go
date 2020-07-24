/*
 * @lc app=leetcode id=797 lang=golang
 *
 * [797] All Paths From Source to Target
 */

package golang

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	paths := [][]int{}
	helpAllPathsSourceTarget(graph, []int{}, 0, &paths)
	return paths
}

func helpAllPathsSourceTarget(graph [][]int, inPath []int, source int, paths *[][]int) {
	curPath := make([]int, len(inPath)+1)
	copy(curPath, inPath)
	curPath[len(inPath)] = source
	if source == len(graph)-1 {
		*paths = append(*paths, curPath)
		return
	}
	for _, next := range graph[source] {
		helpAllPathsSourceTarget(graph, curPath, next, paths)
	}
}

// @lc code=end
