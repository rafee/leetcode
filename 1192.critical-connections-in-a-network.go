/*
 * @lc app=leetcode id=1192 lang=golang
 *
 * [1192] Critical Connections in a Network
 *
 * https://leetcode.com/problems/critical-connections-in-a-network/description/
 *
 * algorithms
 * Hard (48.23%)
 * Likes:    599
 * Dislikes: 57
 * Total Accepted:    29.6K
 * Total Submissions: 61.1K
 * Testcase Example:  '4\n[[0,1],[1,2],[2,0],[1,3]]'
 *
 * There are n servers numbered from 0 to n-1 connected by undirected
 * server-to-server connections forming a network where connections[i] = [a, b]
 * represents a connection between servers a and b. Any server can reach any
 * other server directly or indirectly through the network.
 *
 * A critical connection is a connection that, if removed, will make some
 * server unable to reach some other server.
 *
 * Return all critical connections in the network in any order.
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
 * Output: [[1,3]]
 * Explanation: [[3,1]] is also accepted.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 10^5
 * n-1 <= connections.length <= 10^5
 * connections[i][0] != connections[i][1]
 * There are no repeated connections.
 *
 *
 */

package leetcode

import "fmt"

// @lc code=start
func criticalConnections(n int, connections [][]int) [][]int {
	adjList := make([][]int, n)
	for i := 0; i < n; i++ {
		adjList[i] = make([]int, 0)
	}
	for _, edge := range connections {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		rank[i] = -1
	}
	result := make([][]int, 0)
	highRank := 0
	for _, node := range adjList[0] {
		_, highRank = visitNode(adjList, &result, rank, node, 0, highRank)
	}
	return result
}

func visitNode(graph [][]int, result *[][]int, rank []int, node int, pre int, curRank int) (int, int) {
	lowestRank, highestRank := curRank, curRank
	if rank[node] == -1 {
		fmt.Println(node)
		curRank++
		rank[node] = curRank
		for _, curNode := range graph[node] {
			if curNode == node {
				continue
			}
			lowRank, highRank := visitNode(graph, result, rank, curNode, node, curRank)
			if lowRank > curRank {
				*result = append(*result, []int{node, curNode})
			}
			if lowRank < lowestRank {
				lowestRank = lowRank
			}
			if highRank > highestRank {
				highestRank = highRank
			}
		}
	} else if rank[node] < curRank {
		fmt.Println(node)
		return rank[node], curRank
	}
	return lowestRank, highestRank
}

// @lc code=end
