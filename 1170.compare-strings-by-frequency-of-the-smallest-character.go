/*
 * @lc app=leetcode id=1170 lang=golang
 *
 * [1170] Compare Strings by Frequency of the Smallest Character
 */

package leetcode

import (
	"sort"
)

// @lc code=start
func numSmallerByFrequency(queries []string, words []string) []int {
	count := make([]int, len(words))
	for i, word := range words {
		count[i] = f(word)
	}
	sort.Ints(count)
	answer := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		num := f(queries[i])
		left, right := 0, len(words)-1
		for left < right {
			mid := left + (right-left)/2
			if count[mid] >= num {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		answer[i] = len(words) - left
		// for j := 0; j < len(count); j++ {
		// 	if count[j] > num {
		// 		answer[i] = len(words) - j
		// 		break
		// 	}
		// }
	}
	return answer
}

func f(s string) int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			return freq[i]
		}
	}

	return 0
}

// func min(nums ...int) int {
// 	min := math.MaxInt64
// 	for _, num := range nums {
// 		if num < min {
// 			min = num
// 		}
// 	}
// 	return min
// }

// @lc code=end
