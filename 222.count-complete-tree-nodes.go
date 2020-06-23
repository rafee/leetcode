/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
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
func countNodes(root *TreeNode) int {
	left, right := root, root
	lH, rH := 0, 0
	for left != nil {
		left = left.Left
		lH++
	}
	for right != nil {
		right = right.Right
		rH++
	}

	if lH == rH {
		return (1 << lH) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// @lc code=end
