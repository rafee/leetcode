/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

package leetcode

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	combinationSum3Ref(1, k, n, &result)
	return result
}

func combinationSum3Ref(start int, k int, n int, res *[][]int) {
	if k == 1 && n < 10 && n >= start {
		*res = append(*res, []int{n})
	} else if k > 1 && n > start {
		for i := start; i <= (10 - k); i++ {
			tmp := make([][]int, 0)
			combinationSum3Ref(i+1, k-1, n-i, &tmp)
			for _, item := range tmp {
				*res = append(*res, append(item, i))
			}
		}
	}
}

// @lc code=end
