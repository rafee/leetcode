/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

package golang

// @lc code=start

type state int

const (
	Unexplored state = iota
	Pending
	Explored
)

// func findOrder(numCourses int, prerequisites [][]int) []int {
// 	courseGraph := make([][]int, numCourses)
// 	for _, pre := range prerequisites {
// 		courseGraph[pre[0]] = append(courseGraph[pre[0]], pre[1])
// 	}

// 	starts := make([]int, 0)
// 	for i := 0; i < numCourses; i++ {
// 		if !dependent[i] {
// 			starts = append(starts, i)
// 		}
// 	}

// 	if len(starts) == 0 {
// 		return []int{}
// 	}

// 	fmt.Println(courseGraph)

// 	visited := make([]bool, numCourses)
// 	result := make([]int, 0)
// 	for _, s := range starts {
// 		res := goDeep(courseGraph, visited, s)
// 		if res == nil {
// 			return []int{}
// 		}
// 		result = append(result, res...)
// 	}
// 	return result
// }

// func goDeep(courseGraph [][]int, visited []bool, i int) []int {
// 	result := []int{i}
// 	visited[i] = true
// 	for _, c := range courseGraph[i] {
// 		if visited[c] {
// 			return []int{}
// 		}
// 		res := goDeep(courseGraph, visited, c)
// 		if res == nil {
// 			return []int{}
// 		}
// 		result = append(result, res...)
// 	}
// 	return result
// }

// @lc code=end
