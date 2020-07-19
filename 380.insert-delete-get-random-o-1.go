/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

package golang

import "math/rand"

// @lc code=start
type RandomizedSet struct {
	hashset map[int]int
	nums    []int
}

/** Initialize your data structure here. */
func radomizedConstructor() RandomizedSet {
	randomSet := new(RandomizedSet)
	randomSet.hashset = make(map[int]int)
	randomSet.nums = make([]int, 0)
	return *randomSet
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, found := this.hashset[val]
	if found {
		return false
	}
	this.hashset[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	index, found := this.hashset[val]
	if !found {
		return false
	}
	lastVal := this.nums[len(this.nums)-1]
	this.hashset[lastVal] = index
	this.nums[index] = lastVal
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.hashset, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	random := rand.Intn(len(this.nums))
	return this.nums[random]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
