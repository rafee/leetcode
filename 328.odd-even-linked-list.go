/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
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
func oddEvenList(head *ListNode) *ListNode {
	odd, even := &ListNode{}, &ListNode{}
	oddHead, evenHead := odd, even
	for i := 0; head != nil; i++ {
		if isEven(i) {
			even, head = getNext(even, head)
		} else {
			odd, head = getNext(odd, head)
		}
	}

	odd.Next = nil
	even.Next = oddHead.Next
	return evenHead.Next
}

func getNext(node *ListNode, head *ListNode) (*ListNode, *ListNode) {
	node.Next = head
	return node.Next, head.Next
}

func isEven(num int) bool {
	return (num & 1) == 0
}

// @lc code=end
