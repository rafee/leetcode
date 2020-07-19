/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_helpLevel(t *testing.T) {
	type args struct {
		root  *TreeNode
		depth int
		res   *[][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helpLevel(tt.args.root, tt.args.depth, tt.args.res)
		})
	}
}
