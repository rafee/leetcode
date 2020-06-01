/*
 * @lc app=leetcode id=914 lang=golang
 *
 * [914] X of a Kind in a Deck of Cards
 */

package leetcode

import "math"

// @lc code=start
func hasGroupsSizeX(deck []int) bool {
	numMap := make(map[int]int)
	for _, num := range deck {
		numMap[num]++
	}

	counts := make([]int, 0)
	min := math.MaxInt64
	for _, val := range numMap {
		counts = append(counts, val)
		if val < min {
			min = val
		}
	}

	for {
		flag := true
		for _, count := range counts {
			if count%min != 0 {
				min--
				flag = false
				break
			}
		}
		if flag {
			return min >= 2
		}
	}
}

// @lc code=end
