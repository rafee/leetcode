/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
func kthSmallest(root *TreeNode, k int) int {
	kthRoot, _ := findKSmallest(root, k)
	return kthRoot.Val
}

func findKSmallest(root *TreeNode, k int) (*TreeNode, int) {
	if root == nil {
		return root, 0
	}
	kthRoot, leftCount := findKSmallest(root.Left, k)
	if kthRoot != nil {
		return kthRoot, leftCount
	}
	if leftCount == k-1 {
		return root, k
	}
	kthRoot, rightCount := findKSmallest(root.Right, (k - leftCount - 1))
	return kthRoot, leftCount + rightCount + 1
}

// @lc code=end
