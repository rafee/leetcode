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

package golang

import (
	"reflect"
	"testing"
)

func Test_criticalConnections(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}},
			want: [][]int{{1, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := criticalConnections(tt.args.n, tt.args.connections); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("criticalConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_visitNode(t *testing.T) {
	type args struct {
		graph   [][]int
		result  *[][]int
		rank    []int
		node    int
		pre     int
		curRank int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := visitNode(tt.args.graph, tt.args.result, tt.args.rank, tt.args.node, tt.args.pre, tt.args.curRank)
			if got != tt.want {
				t.Errorf("visitNode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("visitNode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
