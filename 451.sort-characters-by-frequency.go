/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

package leetcode

import (
	"container/heap"
)

// @lc code=start
func frequencySort(s string) string {
	charMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}
	pq := make(freqPQ, 0)
	i:=0
	for key, count := range charMap {
		item := &freqCount{
			char: key,
			freq: count,
			index: i,
		}
		heap.Push(&pq,item)
		i++
	}
	heap.Init(&pq)
	byteArr := make([]byte, 0)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*freqCount)
		if item.freq == 0 {
			break
		}
		for i := 0; i < item.freq; i++ {
			byteArr = append(byteArr, item.char)
		}
	}
	return string(byteArr)
}

// An freqCount is something we manage in a priority queue.
type freqCount struct {
	char byte // The value of the item; arbitrary.
	freq int  // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A freqPQ implements heap.Interface and holds Items.
type freqPQ []*freqCount

func (pq freqPQ) Len() int { return len(pq) }

func (pq freqPQ) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].freq > pq[j].freq
}

func (pq freqPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *freqPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*freqCount)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *freqPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *freqPQ) update(item *freqCount, value byte, priority int) {
	item.char = value
	item.freq = priority
	heap.Fix(pq, item.index)
}

// @lc code=end
