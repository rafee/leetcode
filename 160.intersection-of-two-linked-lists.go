/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		a = nextNode(a, headB)
		b = nextNode(b, headA)
	}
	return a
}

func nextNode(node *ListNode, headB *ListNode) *ListNode {
	if node != nil {
		node = node.Next
	} else {
		node = headB
	}
	return node
}

// @lc code=end
