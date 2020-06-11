/*
 * @lc app=leetcode id=381 lang=golang
 *
 * [381] Insert Delete GetRandom O(1) - Duplicates allowed
 */

package leetcode

import (
	"math/rand"
	"sort"
)

// @lc code=start
type RandomizedCollection struct {
	collection map[int][]int
	nums       []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	col := new(RandomizedCollection)
	col.collection = make(map[int][]int)
	col.nums = make([]int, 0)
	return *col
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	indices, found := this.collection[val]
	if !found {
		indices = make([]int, 0)
	}
	indices = append(indices, len(this.nums))
	this.nums = append(this.nums, val)
	this.collection[val] = indices
	return !found
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	indices, found := this.collection[val]
	if !found {
		return false
	}
	i, indices := popLast(indices)
	var lVal int
	lVal, this.nums = popLast(this.nums)
	if i < len(this.nums) {
		this.nums[i] = lVal
		lIndices := this.collection[lVal]
		_, lIndices = popLast(lIndices)
		lIndices = append(lIndices, i)
		sort.Ints(lIndices)
		this.collection[lVal] = lIndices
	}
	if len(indices) == 0 {
		delete(this.collection, val)
	} else {
		this.collection[val] = indices
	}
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	random := rand.Intn(len(this.nums))
	return this.nums[random]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
