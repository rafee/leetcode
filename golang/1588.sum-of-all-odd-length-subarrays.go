/*
 * @lc app=leetcode id=1588 lang=golang
 *
 * [1588] Sum of All Odd Length Subarrays
 */

package golang

// @lc code=start
func sumOddLengthSubarrays(arr []int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		start, end := len(arr)-i, i+1
		total := start * end
		odd := (total + 1) / 2
		result += (arr[i] * odd)
	}
	return result
}

// @lc code=end
