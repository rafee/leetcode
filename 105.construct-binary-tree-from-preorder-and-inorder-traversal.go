/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */

package leetcode

import "math"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	root, _, _ := makeTree(preorder, inorder, math.MaxInt64)
	return root
}

func makeTree(preorder []int, inorder []int, limit int) (*TreeNode, []int, []int) {
	root := &TreeNode{}
	root.Val = preorder[0]
	root.Left, preorder, inorder = makeTree(preorder[1:], inorder, preorder[0])
	root.Right, preorder, inorder = makeTree(preorder, inorder, limit)
	return root, preorder, inorder
}

// @lc code=end
