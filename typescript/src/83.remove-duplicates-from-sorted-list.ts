/*
 * @lc app=leetcode id=83 lang=typescript
 *
 * [83] Remove Duplicates from Sorted List
 */


// @lc code=start

function deleteDuplicates(head: ListNode | null): ListNode | null {
    if (head === null) {
        return null
    }
    let curNode = head
    while (curNode.next !== null) {
        if (curNode.val === curNode.next.val) {
            curNode.next = curNode.next.next
        } else {
            curNode = curNode.next
        }
    }
    return head
};
// @lc code=end

