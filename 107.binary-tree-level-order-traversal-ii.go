/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	helpLevel(root, 0, &res)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

func helpLevel(root *TreeNode, depth int, res *[][]int) {
	if depth == len(*res) {
		*res = append(*res, []int{})
	}
	(*res)[depth] = append((*res)[depth], root.Val)
	if root.Left != nil {
		helpLevel(root.Left, depth+1, res)
	}
	if root.Right != nil {
		helpLevel(root.Right, depth+1, res)
	}
}

// @lc code=end
