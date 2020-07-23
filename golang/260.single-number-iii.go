/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 */

package golang

// @lc code=start
func singleNumberIII(nums []int) []int {
	joint := 0
	for _, num := range nums {
		joint ^= num
	}

	single := joint & -joint

	one := 0
	for _, num := range nums {
		if num&single == 0 {
			one ^= num
		}
	}
	return []int{one, joint ^ one}
}

// @lc code=end
