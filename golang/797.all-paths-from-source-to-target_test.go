/*
 * @lc app=leetcode id=797 lang=golang
 *
 * [797] All Paths From Source to Target
 */

package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{[][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}},
			want: [][]int{{0, 3, 6, 7}, {0, 3, 4, 7}, {0, 3, 4, 6, 7}, {0, 3, 4, 5, 6, 7}, {0, 1, 4, 7}, {0, 1, 4, 6, 7}, {0, 1, 4, 5, 6, 7}, {0, 1, 6, 7}, {0, 1, 7}, {0, 1, 2, 4, 7}, {0, 1, 2, 4, 6, 7}, {0, 1, 2, 4, 5, 6, 7}, {0, 1, 2, 6, 7}, {0, 1, 2, 3, 6, 7}, {0, 1, 2, 3, 4, 7}, {0, 1, 2, 3, 4, 6, 7}, {0, 1, 2, 3, 4, 5, 6, 7}, {0, 1, 5, 6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, allPathsSourceTarget(tt.args.graph))
		})
	}
}

func Test_helpAllPathsSourceTarget(t *testing.T) {
	type args struct {
		graph  [][]int
		inPath []int
		source int
		paths  *[][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helpAllPathsSourceTarget(tt.args.graph, tt.args.inPath, tt.args.source, tt.args.paths)
		})
	}
}
