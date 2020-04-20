/*
 * @lc app=leetcode id=1008 lang=golang
 *
 * [1008] Construct Binary Search Tree from Preorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/description/
 *
 * algorithms
 * Medium (74.42%)
 * Likes:    694
 * Dislikes: 23
 * Total Accepted:    51.1K
 * Total Submissions: 67K
 * Testcase Example:  '[8,5,1,7,10,12]'
 *
 * Return the root node of a binary search tree that matches the given preorder
 * traversal.
 *
 * (Recall that a binary search tree is a binary tree where for every node, any
 * descendant of node.left has a value < node.val, and any descendant of
 * node.right has a value > node.val.  Also recall that a preorder traversal
 * displays the value of the node first, then traverses node.left, then
 * traverses node.right.)
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [8,5,1,7,10,12]
 * Output: [8,5,10,1,7,null,12]
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= preorder.length <= 100
 * The values of preorder are distinct.
 *
 *
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_bstFromPreorder(t *testing.T) {
	type args struct {
		preorder []int
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
			if got := bstFromPreorder(tt.args.preorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bstWithArr(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name  string
		args  args
		want  *TreeNode
		want1 []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := bstWithArr(tt.args.nums, tt.args.limit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstWithArr() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("bstWithArr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
