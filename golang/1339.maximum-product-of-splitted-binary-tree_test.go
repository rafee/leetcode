/*
 * @lc app=leetcode id=1339 lang=golang
 *
 * [1339] Maximum Product of Splitted Binary Tree
 */

package golang

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := maxProduct(tt.args.root); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
