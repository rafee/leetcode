/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_inpost_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		in_begin  int
		in_end    int
		postorder []int
		root_pos  int
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
			if got := inpost_buildTree(tt.args.inorder, tt.args.in_begin, tt.args.in_end, tt.args.postorder, tt.args.root_pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inpost_buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
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
			if got := buildTree(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
