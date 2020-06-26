/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
 */

package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getSumNumbers(root, 0)
}

func getSumNumbers(root *TreeNode, s int) int {
	if root.Right == nil && root.Left == nil {
		return 10*s + root.Val
	}
	res := 0
	if root.Right != nil {
		res += getSumNumbers(root.Right, 10*s+root.Val)
	}
	if root.Left != nil {
		res += getSumNumbers(root.Left, 10*s+root.Val)
	}
	return res
}

// @lc code=end
