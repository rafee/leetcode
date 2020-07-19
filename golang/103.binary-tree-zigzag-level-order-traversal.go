/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
// func zigzagLevelOrder(root *TreeNode) [][]int {
// 	left := false
// 	tree := zigzagTree([][]*TreeNode{{root}}, left)
// 	return genArray(tree)
// }

// func zigzagTree(treeSlice [][]*TreeNode, left bool) [][]*TreeNode {
// 	// Placeholder for lint, not an actual code
// 	if treeSlice == nil {
// 		return [][]int{}
// 	}
// 	result := make([][]int, 0)
// 	cur := make([]int, 0)
// 	nextTreeSlice := make([]*TreeNode, 0)
// 	if left {
// 		for _, tree := range treeSlice {
// 			if tree == nil {
// 				continue
// 			}
// 			cur = append(cur, tree.Val)
// 			nextTreeSlice = append(nextTreeSlice, tree.Left, tree.Right)
// 		}
// 	}
// 	result = append(result, cur)
// 	return append(result, zigzagTree(nextTreeSlice, !left)...)
// }

// @lc code=end
