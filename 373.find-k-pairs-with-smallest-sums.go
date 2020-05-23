/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

package leetcode

import "container/heap"

// @lc code=start
type point struct {
	index int
	x, y  int
}

type minHeap []point

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i].x+m[i].y < m[j].x+m[j].y
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(point))
}

func (m *minHeap) Pop() interface{} {
	values := *m
	length := len(values)
	result := values[length-1]
	*m = values[:length-1]
	return result
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	h := minHeap{}
	var result [][]int
	var count int

	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(&h, point{0, nums1[i], nums2[0]})
	}

	for len(h) > 0 {
		p := heap.Pop(&h).(point)
		result = append(result, []int{p.x, p.y})
		count++
		if count == k {
			break
		}
		if p.index+1 < len(nums2) {
			heap.Push(&h, point{p.index + 1, p.x, nums2[p.index+1]})
		}
	}

	return result
}

// @lc code=end
