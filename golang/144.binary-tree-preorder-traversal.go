/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */

package golang

// @lc code=start
func preorderTraversal(root *TreeNode) []int {
	return helpPreorderTraversalRecurive(root, []int{})
}

func helpPreorderTraversalRecurive(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	result = helpPreorderTraversalRecurive(root.Left, append(result, root.Val))
	result = helpPreorderTraversalRecurive(root.Right, result)
	return result
}

func helpPreorderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	var s stk
	s.push(root)
	for !s.isEmpty() {
		cur:=s.pop()
		result=append(result, cur.Val)
		if cur.Left!=nil{
			s.push(cur.Left)
		}
	}
	return result
}

type stk []*TreeNode

func (s *stk) push(node *TreeNode) {
	*s = append(*s, node)
}
func (s *stk) pop() *TreeNode {
	i := len(*s) - 1
	result := (*s)[i]
	*s=(*s)[:i]
	return result
}

func (s *stk) isEmpty() bool {
	return len(*s) == 0
}

// @lc code=end
