/*
 * @lc app=leetcode id=528 lang=golang
 *
 * [528] Random Pick with Weight
 */

package leetcode

import "math/rand"

// @lc code=start
type Solution struct {
	cumSum []int
	limit  int
}

func Constructor(w []int) Solution {
	sol := new(Solution)
	sol.cumSum = make([]int, len(w))
	val := 0
	for i, num := range w {
		val += num
		sol.cumSum[i] = val
	}
	sol.limit = val
	return *sol
}

func (this *Solution) PickIndex() int {
	val := rand.Intn(this.limit) + 1
	return this.bSearch(val)
}

func (this *Solution) bSearch(target int) int {
	left, right := 0, len(this.cumSum)
	for left < right {
		mid := left + (right-left)/2
		if this.cumSum[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// @lc code=end
