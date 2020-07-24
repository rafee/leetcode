/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 *
 * https://leetcode.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (44.57%)
 * Likes:    1426
 * Dislikes: 61
 * Total Accepted:    259K
 * Total Submissions: 566.4K
 * Testcase Example:  '[1,2,2]'
 *
 * Given a collection of integers that might contain duplicates, nums, return
 * all possible subsets (the power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: [1,2,2]
 * Output:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
 *
 */

package golang

import (
	"sort"
)

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{{}}
	var cur []int

	sort.Ints(nums)
	doSubsetsWithDup(nums, cur, &res)
	return res
}

func doSubsetsWithDup(nums []int, cur []int, res *[][]int) {
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		cur = append(cur, nums[i])
		newRes := make([]int, len(cur))
		copy(newRes, cur)
		*res = append(*res, newRes)

		doSubsetsWithDup(nums[i+1:], cur, res)
		cur = cur[:len(cur)-1]
	}
}

// @lc code=end
