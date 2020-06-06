/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
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
func sortList(head *ListNode) *ListNode {
	head, _ = sortListHelper(head)
	return head
}

func sortListHelper(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	pivot, cur, curHead, pre := head, head.Next, &ListNode{}, head
	for cur != nil {
		if cur.Val < pivot.Val {
			pre.Next = cur.Next
			cur.Next = curHead
			curHead = cur
		} else {
			pre.Next = cur
			pre = pre.Next
		}
		cur = pre.Next
	}
	startRight, endRight := sortListHelper(pivot.Next)
	startLeft, endLeft := sortListHelper(curHead)
	endLeft.Next = pivot
	pivot.Next = startRight
	return startLeft, endRight
}

// @lc code=end
