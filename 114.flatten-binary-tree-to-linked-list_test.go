/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_helpFlatten(t *testing.T) {
	type args struct {
		root  *TreeNode
		right *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helpFlatten(tt.args.root, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helpFlatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flattenBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flattenBinaryTree(tt.args.root)
		})
	}
}
