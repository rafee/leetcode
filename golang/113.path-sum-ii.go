/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 *
 * https://leetcode.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (43.39%)
 * Likes:    1472
 * Dislikes: 51
 * Total Accepted:    306.5K
 * Total Submissions: 680.9K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * Given a binary tree and a sum, find all root-to-leaf paths where each path's
 * sum equals the given sum.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given the below binary tree and sum = 22,
 *
 *
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \    / \
 * 7    2  5   1
 *
 *
 * Return:
 *
 *
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
 * ]
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
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	if root.Val == sum && root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	left := pathSum(root.Left, sum-root.Val)
	right := pathSum(root.Right, sum-root.Val)

	if left == nil && right == nil {
		return nil
	}

	cur := []int{root.Val}
	res := [][]int{}

	for _, arr := range left {
		res = append(res, append(cur, arr...))
	}

	for _, arr := range right {
		res = append(res, append(cur, arr...))
	}

	return res
}

// @lc code=end
