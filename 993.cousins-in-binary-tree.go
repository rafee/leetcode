/*
 * @lc app=leetcode id=993 lang=golang
 *
 * [993] Cousins in Binary Tree
 *
 * https://leetcode.com/problems/cousins-in-binary-tree/description/
 *
 * algorithms
 * Easy (51.97%)
 * Likes:    599
 * Dislikes: 33
 * Total Accepted:    61.7K
 * Total Submissions: 116.6K
 * Testcase Example:  '[1,2,3,4]\n4\n3'
 *
 * In a binary tree, the root node is at depth 0, and children of each depth k
 * node are at depth k+1.
 *
 * Two nodes of a binary tree are cousins if they have the same depth, but have
 * different parents.
 *
 * We are given the root of a binary tree with unique values, and the values x
 * and y of two different nodes in the tree.
 *
 * Return true if and only if the nodes corresponding to the values x and y are
 * cousins.
 *
 *
 *
 * Example 1:
 *
 *
 *
 * Input: root = [1,2,3,4], x = 4, y = 3
 * Output: false
 *
 *
 *
 * Example 2:
 *
 *
 *
 * Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 *
 *
 * Input: root = [1,2,3,null,4], x = 2, y = 3
 * Output: false
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of nodes in the tree will be between 2 and 100.
 * Each node has a unique integer value from 1 to 100.
 *
 *
 *
 *
 *
 *
 *
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
func isCousins(root *TreeNode, x int, y int) bool {
	if root.Val == x || root.Val == y {
		return false
	}
	return isCousinsHelper([]*TreeNode{root}, x, y)
}

func isCousinsHelper(nodes []*TreeNode, x int, y int) bool {
	if len(nodes) == 0 {
		return false
	}
	xp, yp, next := -1, -1, []*TreeNode{}
	for _, node := range nodes {
		if node.Left != nil {
			if node.Left.Val == x {
				xp = node.Val
			} else if node.Left.Val == y {
				yp = node.Val
			}
			next = append(next, node.Left)
		}
		if node.Right != nil {
			if node.Right.Val == x {
				xp = node.Val
			} else if node.Right.Val == y {
				yp = node.Val
			}
			next = append(next, node.Right)
		}
	}
	if xp != -1 || yp != -1 {
		return xp != yp && xp != -1 && yp != -1
	}
	return isCousinsHelper(next, x, y)
}

// @lc code=end
