/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_bstConstructor(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want BSTIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstConstructor(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstConstructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTIterator_Next(t *testing.T) {
	tests := []struct {
		name string
		this *BSTIterator
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Next(); got != tt.want {
				t.Errorf("BSTIterator.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTIterator_HasNext(t *testing.T) {
	tests := []struct {
		name string
		this *BSTIterator
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.HasNext(); got != tt.want {
				t.Errorf("BSTIterator.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
