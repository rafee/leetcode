/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_helpPreorderTraversalRecurive(t *testing.T) {
	type args struct {
		root   *TreeNode
		result []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helpPreorderTraversalRecurive(tt.args.root, tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helpPreorderTraversalRecurive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_helpPreorderTraversalIterative(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helpPreorderTraversalIterative(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helpPreorderTraversalIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stk_push(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		s    *stk
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.push(tt.args.node)
		})
	}
}

func Test_stk_pop(t *testing.T) {
	tests := []struct {
		name string
		s    *stk
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stk.pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stk_isEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    *stk
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.isEmpty(); got != tt.want {
				t.Errorf("stk.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
