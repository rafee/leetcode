/*
 * @lc app=leetcode id=1822 lang=golang
 *
 * [1822] Sign of the Product of an Array
 */

package golang

// @lc code=start
func arraySign(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			i++
		}
	}
	if i%2 == 1 {
		return -1
	}
	return 1
}

// @lc code=end
