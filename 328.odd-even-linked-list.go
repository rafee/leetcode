/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
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
// func oddEvenList(head *ListNode) *ListNode {
// 	odd, even := head, &ListNode{}
// 	oddIterator, evenIterator := &ListNode{}, &ListNode{}
// 	for i := 0; head != nil; i++ {
// 		if i%2 == 0 {
// 			odd = head
// 			head = head.Next
// 			odd = odd.Next
// 		} else {
// 			even = head
// 			head = head.Next
// 		}
// 	}
// 	odd.Next = even
// 	return odd
// }

// @lc code=end
