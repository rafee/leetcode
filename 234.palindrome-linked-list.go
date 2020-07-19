/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

package golang

// @lc code=start
func isPalindrome(head *ListNode) bool {
	reverseList := &ListNode{}
	for tmp, start := (&ListNode{}), head; start != nil; start = start.Next {
		tmp = start.Next
		start.Next = reverseList
		reverseList, start = start, tmp
	}
	for reverseList != nil {
		if reverseList.Val != head.Val {
			return false
		}
		reverseList = reverseList.Next
		head = head.Next
	}
	return true
}

// @lc code=end
