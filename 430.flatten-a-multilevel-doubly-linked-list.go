/*
 * @lc app=leetcode id=430 lang=golang
 *
 * [430] Flatten a Multilevel Doubly Linked List
 */

package leetcode

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	head, _ := flattenNodeHelper(root)
	return head
}

func flattenNodeHelper(root *Node) (*Node, *Node) {
	head := root
	for root.Next != nil {
		if root.Child == nil {
			root = root.Next
		} else {
			next := root.Next
			root = fixChild(root)
			root.Next, next.Prev = next, root
		}
	}
	// When the last node has child. Special Case
	if root.Child != nil {
		root = fixChild(root)
	}
	return head, root
}

func fixChild(root *Node) *Node {
	child, end := flattenNodeHelper(root.Child)
	root.Child = nil
	root.Next, child.Prev = child, root
	root = end
	return root
}

// @lc code=end
