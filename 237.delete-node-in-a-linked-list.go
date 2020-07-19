/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */

package golang

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	pre := node
	for node.Next != nil {
		node.Val = node.Next.Val
		pre = node
		node = node.Next
	}
	pre.Next = nil
}

// @lc code=end
