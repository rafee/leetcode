package leetcode

import "strconv"

// ListNode defines a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// newList creates a new list from list of ints
func newList(nums []int) *ListNode {
	listNode := &ListNode{}
	original := listNode
	for _, v := range nums {
		temp := &ListNode{}
		temp.Val = v
		listNode.Next = temp
		listNode = listNode.Next
	}
	return original.Next
}

// printList prints a ListNode for debugging purpose
func printList(l *ListNode) (s string) {
	s = ""
	for l != nil {
		s += (strconv.Itoa(l.Val) + " ")
		l = l.Next
	}
	return s
}

// A pqLinkedList implements heap.Interface and holds Items.
type pqLinkedList []*ListNode

func (pq pqLinkedList) Len() int { return len(pq) }

func (pq pqLinkedList) Less(i, j int) bool {
	// Get the lowest Node from queue
	return pq[i].Val < pq[j].Val
}

func (pq pqLinkedList) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pqLinkedList) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *pqLinkedList) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
