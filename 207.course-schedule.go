/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

package leetcode

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	courseGraph := make([][]int, numCourses)

	for i := 0; i < numCourses; i++ {
		courseGraph[i] = make([]int, 0)
	}

	for _, preReq := range prerequisites {
		courseGraph[preReq[1]] = append(courseGraph[preReq[1]], preReq[0])
	}

	visited := make(map[int]bool)

	for i := 0; i < numCourses; i++ {
		if !nextCourse(courseGraph, visited, i) {
			return false
		}
	}
	return true
}

func nextCourse(graph [][]int, visited map[int]bool, node int) bool {
	_, found := visited[node]
	if found {
		return visited[node]
	}
	visited[node] = false
	for _, nextNode := range graph[node] {
		if !nextCourse(graph, visited, nextNode){
			return false
		}
	}
	visited[node] = true
	return true
}

// @lc code=end
