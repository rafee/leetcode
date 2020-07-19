/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */

package golang

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flattenBinaryTree(root *TreeNode) {
	helpFlatten(root, nil)
}

func helpFlatten(root *TreeNode, right *TreeNode) *TreeNode {
	if root == nil {
		return right
	}
	root.Right = helpFlatten(root.Right, right)
	root.Right = helpFlatten(root.Left, root.Right)
	root.Left = nil
	return root
}

// @lc code=end
