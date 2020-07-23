/* eslint-disable no-param-reassign */
/*
 * @lc app=leetcode id=203 lang=typescript
 *
 * [203] Remove Linked List Elements
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function removeElements(head: ListNode | null, val: number): ListNode | null {
  while (head !== null && head.val === val) {
    head = head.next
  }
  if (head === null) {
    return head
  }
  let curNode = head
  while (curNode.next !== null) {
    if (curNode.next.val === val) {
      curNode.next = curNode.next.next
    } else {
      curNode = curNode.next
    }
  }
  return head
}
// @lc code=end
