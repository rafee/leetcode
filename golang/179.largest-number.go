/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */

package golang

import (
	"sort"
	"strconv"
	"strings"
)

// @lc code=start
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	if strs[0][0] == '0' {
		return "0"
	}
	return strings.Join(strs, "")
}

// @lc code=end
