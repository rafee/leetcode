/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getSumNumbers(root, 0)
}

func getSumNumbers(root *TreeNode, s int) int {
	curRootSum := 10*s + root.Val
	if root.Right == nil && root.Left == nil {
		return curRootSum
	}
	sum := 0
	if root.Right != nil {
		sum += getSumNumbers(root.Right, curRootSum)
	}
	if root.Left != nil {
		sum += getSumNumbers(root.Left, curRootSum)
	}
	return sum
}

// @lc code=end
