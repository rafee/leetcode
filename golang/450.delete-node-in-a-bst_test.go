/*
 * @lc app=leetcode id=450 lang=golang
 *
 * [450] Delete Node in a BST
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
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
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinNode(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := findMinNode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
