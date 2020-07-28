/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_buildTreePreIn(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
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
			if got := buildTreePreIn(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTreePreIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeTreePreIn(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
		limit    int
	}
	tests := []struct {
		name  string
		args  args
		want  *TreeNode
		want1 []int
		want2 []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := makeTreePreIn(tt.args.preorder, tt.args.inorder, tt.args.limit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeTreePreIn() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("makeTreePreIn() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("makeTreePreIn() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
