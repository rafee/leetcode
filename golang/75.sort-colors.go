/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

package golang

// @lc code=start
func sortColors(nums []int) {
	red, white, blue := 0, 0, len(nums)-1
	for white <= blue {
		if nums[white] == 0 {
			nums[red], nums[white] = nums[white], nums[red]
			red++
			white++
		} else if nums[white] == 1 {
			white++
		} else {
			nums[blue], nums[white] = nums[white], nums[blue]
			blue--
		}
	}
}

// @lc code=end
