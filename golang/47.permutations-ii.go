/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

package golang

import (
	"log"
	"sort"
)

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{{}}
	for i, num := range nums {
		// New slice to hold results from current step
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		curResult := [][]int{}
		// Get all the permutations in current result
		for _, per := range result {
			for i := 0; i < len(per)+1; i++ {
				// Make a slice one element bigger
				item := make([]int, len(per)+1)
				item[i] = num
				copy(item[:i], per[:i])
				copy(item[i+1:], per[i:])
				curResult = append(curResult, item)
			}
		}
		// Make current result the new result
		result = curResult
	}
	return result
}

func genPermutationsII(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		next := append(nums[0:i], nums[i+1:]...)
		log.Println(next)
		nextRes := genPermutationsII(next)
		comb := []int{nums[i]}
		for _, per := range nextRes {
			result = append(result, append(comb, per...))
		}
	}
	return result
}

// @lc code=end
