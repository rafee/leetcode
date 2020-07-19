/*
 * @lc app=leetcode id=987 lang=golang
 *
 * [987] Vertical Order Traversal of a Binary Tree
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_verticalTraversal(t *testing.T) {
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
			if got := verticalTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_helpVertical(t *testing.T) {
	type args struct {
		root  *TreeNode
		left  *[][]int
		right *[][]int
		pos   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helpVertical(tt.args.root, tt.args.left, tt.args.right, tt.args.pos)
		})
	}
}
