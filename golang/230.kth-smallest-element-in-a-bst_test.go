/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name  string
		args  args
		want  *TreeNode
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findKSmallest(tt.args.root, tt.args.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findKSmallest() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findKSmallest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
