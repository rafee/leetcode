/*
 * @lc app=leetcode id=948 lang=golang
 *
 * [948] Bag of Tokens
 */

package golang

import "sort"

// @lc code=start
func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	sums := make([]int, len(tokens)+1)
	sums[0] = 0
	for i := 0; i < len(tokens); i++ {
		sums[i+1] = sums[i] + tokens[i]
	}
	if len(tokens) == 0 || tokens[0] > P {
		return 0
	}
	for left, right := 0, len(tokens)-1; left <= right; {
		if sums[right+1]-sums[left] <= P {
			return right - left + 1
		} else if sums[right]-sums[left] <= P {
			return right - left
		}
		P += (tokens[right] - tokens[left])
		left++
		right--
	}
	return 0
}

// @lc code=end
