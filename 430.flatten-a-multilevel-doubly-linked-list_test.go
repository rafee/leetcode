/*
 * @lc app=leetcode id=430 lang=golang
 *
 * [430] Flatten a Multilevel Doubly Linked List
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatten(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flattenNodeHelper(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name  string
		args  args
		want  *Node
		want1 *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := flattenNodeHelper(tt.args.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenNodeHelper() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("flattenNodeHelper() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_fixChild(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fixChild(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fixChild() = %v, want %v", got, tt.want)
			}
		})
	}
}
