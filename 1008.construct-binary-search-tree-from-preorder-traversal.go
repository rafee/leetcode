/*
 * @lc app=leetcode id=1008 lang=golang
 *
 * [1008] Construct Binary Search Tree from Preorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/description/
 *
 * algorithms
 * Medium (74.42%)
 * Likes:    694
 * Dislikes: 23
 * Total Accepted:    51.1K
 * Total Submissions: 67K
 * Testcase Example:  '[8,5,1,7,10,12]'
 *
 * Return the root node of a binary search tree that matches the given preorder
 * traversal.
 *
 * (Recall that a binary search tree is a binary tree where for every node, any
 * descendant of node.left has a value < node.val, and any descendant of
 * node.right has a value > node.val.  Also recall that a preorder traversal
 * displays the value of the node first, then traverses node.left, then
 * traverses node.right.)
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [8,5,1,7,10,12]
 * Output: [8,5,10,1,7,null,12]
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= preorder.length <= 100
 * The values of preorder are distinct.
 *
 *
 */

package leetcode

import "math"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	root, _ := bstWithArr(preorder, math.MaxInt64)
	return root
}

func bstWithArr(nums []int, limit int) (*TreeNode, []int) {
	if len(nums) == 0 || nums[0] > limit {
		return nil, nums
	}
	root := &TreeNode{nums[0], nil, nil}
	if len(nums) == 1 {
		return root, []int{}
	}
	var numsRight, numsTop []int
	root.Left, numsRight = bstWithArr(nums[1:], nums[0])
	root.Right, numsTop = bstWithArr(numsRight, limit)
	return root, numsTop
}

// @lc code=end
