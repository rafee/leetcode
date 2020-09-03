/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 */

package golang

// @lc code=start

func inpostBuildTree(inorder []int, inBegin int, inEnd int, postOrder []int, rootPos int) *TreeNode {
	if rootPos < 0 || inBegin > inEnd {
		return nil
	}
	//find root's position in inorder
	var i int = inBegin
	for i < inEnd && inorder[i] != postOrder[rootPos] {
		i++
	}
	var node TreeNode
	node.Val = postOrder[rootPos]
	node.Left = inpostBuildTree(inorder, inBegin, i-1, postOrder, rootPos-inEnd+i-1)
	node.Right = inpostBuildTree(inorder, i+1, inEnd, postOrder, rootPos-1)
	return &node
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	lin := len(inorder)
	lpost := len(postorder)
	if lin == 0 || lpost == 0 || lin != lpost {
		return nil
	}
	return inpostBuildTree(inorder, 0, lin-1, postorder, lpost-1)
}

// @lc code=end
