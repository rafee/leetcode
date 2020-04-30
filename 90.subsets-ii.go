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

package leetcode

import (
	"sort"
)

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return genSubsets(nums)
}

func genSubsets(nums []int) [][]int {
	numLen := len(nums)
	if numLen == 0 {
		return [][]int{{}}
	}

	dups, nums := genDup(nums)
	solution := [][]int{}
	res := genSubsets(nums)
	solution = append(solution, res...)
	for _, dup := range dups {
		for _, set := range res {
			tmp := append(dup, set...)
			solution = append(solution, tmp)
		}
	}
	return solution
}

func genDup(nums []int) ([][]int, []int) {
	dups := [][]int{{nums[0]}}
	numLen := len(nums)
	for i := 1; i < numLen; i++ {
		if nums[i] != nums[0] {
			return dups, nums[i:]
		}
		tmp := []int{nums[0]}
		for j := 0; j < i; j++ {
			tmp = append(tmp, nums[0])
		}
		dups = append(dups, tmp)
	}
	return dups, []int{}
}

// @lc code=end
