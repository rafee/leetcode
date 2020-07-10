/*
 * @lc app=leetcode id=662 lang=golang
 *
 * [662] Maximum Width of Binary Tree
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
func widthOfBinaryTree(root *TreeNode) int {
	width := 0
	getWidthHelperBFS([]*TreeNode{root}, 0, []int{0}, &width)
	return width
}

func getWidthHelperBFS(nodes []*TreeNode, level int, order []int, width *int) {
	if len(nodes) == 0 {
		return
	}
	*width = max(*width, order[len(nodes)-1]-order[0]+1)
	nextLevel := []*TreeNode{}
	nextOrder := []int{}
	for i, v := range nodes {
		if v.Left != nil {
			nextLevel = append(nextLevel, v.Left)
			nextOrder = append(nextOrder, order[i]*2)
		}
		if v.Right != nil {
			nextLevel = append(nextLevel, v.Right)
			nextOrder = append(nextOrder, order[i]*2+1)
		}
	}
	getWidthHelperBFS(nextLevel, level+1, nextOrder, width)
}

// @lc code=end
