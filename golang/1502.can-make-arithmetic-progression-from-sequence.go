/*
 * @lc app=leetcode id=1502 lang=golang
 *
 * [1502] Can Make Arithmetic Progression From Sequence
 */

package golang

import "sort"

// @lc code=start
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[0] - arr[1]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i]-arr[i+1] != diff {
			return false
		}
	}
	return true
}

// @lc code=end
