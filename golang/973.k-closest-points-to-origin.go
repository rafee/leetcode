/*
 * @lc app=leetcode id=973 lang=golang
 *
 * [973] K Closest Points to Origin
 */

package golang

import "container/heap"

// @lc code=start
func kClosest(points [][]int, K int) [][]int {
	return kClosestSelect(points, K)
}

func kClosestSelect(points [][]int, K int) [][]int {
	quickSelect(points, 0, len(points)-1, K)
	return points[:K]
}

func quickSelect(points [][]int, start, end, K int) {
	pivot := start
	for i := start; i < end; i++ {
		distCur := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		distEnd := points[end][0]*points[end][0] + points[end][1]*points[end][1]
		if distCur <= distEnd {
			swapPoints(pivot, i, points)
			pivot++
		}
	}

	swapPoints(pivot, end, points)
	if pivot+1 == K {
		return
	} else if pivot+1 > K {
		quickSelect(points, start, pivot-1, K)
	} else {
		quickSelect(points, pivot+1, end, K)
	}
}

func swapPoints(i int, pivot int, points [][]int) {
	points[i], points[pivot] = points[pivot], points[i]
}

func kClosestHeap(points [][]int, K int) [][]int {
	if len(points) == 0 || K == 0 {
		return [][]int{}
	}

	h := closePoints{}
	result := make([][]int, 0)

	for _, point := range points[:K] {
		heap.Push(&h, closePoint{x: point[0], y: point[1], dist: point[0]*point[0] + point[1]*point[1]})
	}

	for i := K; i < len(points); i++ {
		heap.Push(&h, closePoint{x: points[i][0], y: points[i][1], dist: points[i][0]*points[i][0] + points[i][1]*points[i][1]})
		heap.Pop(&h)
	}

	for len(h) > 0 {
		point := heap.Pop(&h).(closePoint)
		result = append(result, []int{point.x, point.y})
	}

	return result
}

type closePoint struct {
	x, y int
	dist int
}

type closePoints []closePoint

func (c closePoints) Len() int {
	return len(c)
}

func (c closePoints) Less(i, j int) bool {
	return c[i].dist > c[j].dist
}

func (c closePoints) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *closePoints) Push(p interface{}) {
	*c = append(*c, p.(closePoint))
}

func (c *closePoints) Pop() interface{} {
	arr := *c
	length := len(arr)
	ret := arr[length-1]
	*c = arr[:length-1]
	return ret
}

// @lc code=end
