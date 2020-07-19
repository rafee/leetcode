/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 */

package golang

import "sort"

// @lc code=start
func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	combs := make([]int, target+1)
	combs[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i < num {
				break
			}
			combs[i] += combs[i-num]
		}
	}
	return combs[target]
}

// @lc code=end
