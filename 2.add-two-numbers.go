/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (32.37%)
 * Likes:    6592
 * Dislikes: 1714
 * Total Accepted:    1.1M
 * Total Submissions: 3.5M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Example:
 *
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 *
 *
 */

package leetcode

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode is a data structure for Linked-List Data type
// @lc code=start

func addTwoNumbersLSBFirst(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	carry := 0
	value := 0
	for {
		if l2 == nil && l1 != nil {
			value = l1.Val + carry
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			value = l2.Val + carry
			l2 = l2.Next
		} else if l1 == nil && l2 == nil {
			break
		} else {
			value = l1.Val + l2.Val + carry
			l1 = l1.Next
			l2 = l2.Next
		}

		cur.Next = &ListNode{value % 10, nil}
		carry = value / 10
		cur = cur.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{carry, nil}
	}

	return result.Next
}

// @lc code=end
