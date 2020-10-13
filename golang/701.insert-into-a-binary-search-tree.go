/*
 * @lc app=leetcode id=701 lang=golang
 *
 * [701] Insert into a Binary Search Tree
 */

package golang

// @lc code=start

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return insertIntoBSTIterative(root, val)
}

func insertIntoBSTRecursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	} else if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func insertIntoBSTIterative(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	for iter := root; ; {
		if iter.Val > val {
			if iter.Left == nil {
				iter.Left = &TreeNode{Val: val}
				break
			} else {
				iter = iter.Left
			}
		} else {
			if iter.Right == nil {
				iter.Right = &TreeNode{Val: val}
				break
			} else {
				iter = iter.Right
			}
		}
	}
	return root
}

// @lc code=end
