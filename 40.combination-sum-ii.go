/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (44.64%)
 * Likes:    1494
 * Dislikes: 56
 * Total Accepted:    302.5K
 * Total Submissions: 651.3K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * Given a collection of candidate numbers (candidates) and a target number
 * (target), find all unique combinations in candidates where the candidate
 * numbers sums to target.
 *
 * Each number in candidates may only be used once in the combination.
 *
 * Note:
 *
 *
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [10,1,2,7,6,1,5], target = 8,
 * A solution set is:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,5,2,1,2], target = 5,
 * A solution set is:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 */

package leetcode

import (
	"sort"
)

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	return genCombinations(candidates, target)
}

func genCombinations(candidates []int, target int) [][]int {
	solution := make([][]int, 0)
	for i, candidate := range candidates {
		if i != 0 && candidates[i] == candidates[i-1] {
			continue
		}
		if target == candidate {
			return append(solution, []int{target})
		} else if target < candidate {
			return solution
		}
		tmp := genCombinations(candidates[i+1:], target-candidate)
		for _, comb := range tmp {
			solution = append(solution, append(comb, candidate))
		}
	}
	return solution
}

// @lc code=end
