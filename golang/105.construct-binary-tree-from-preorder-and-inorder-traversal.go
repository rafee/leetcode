/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */

package golang

import "math"

// @lc code=start
func buildTreePreIn(preorder []int, inorder []int) *TreeNode {
	root, _, _ := makeTreePreIn(preorder, inorder, math.MaxInt64)
	return root
}

func makeTreePreIn(preorder []int, inorder []int, limit int) (*TreeNode, []int, []int) {
	root := &TreeNode{}
	// Placeholder, not real code
	if len(preorder) == 0 {
		return root, preorder, inorder
	}
	root.Val = preorder[0]
	root.Left, preorder, inorder = makeTreePreIn(preorder[1:], inorder, preorder[0])
	root.Right, preorder, inorder = makeTreePreIn(preorder, inorder, limit)
	return root, preorder, inorder
}

// @lc code=end
