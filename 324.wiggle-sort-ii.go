/*
 * @lc app=leetcode id=324 lang=golang
 *
 * [324] Wiggle Sort II
 */

package leetcode

// @lc code=start
func wiggleSort(nums []int) {
	pre := nums[1]
	diff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		curDiff := nums[i] - pre
		if (diff > 0 && curDiff < 0) || (diff < 0 && curDiff > 0) {
			diff = curDiff
			pre = nums[i]
		} else {
			if diff < 0 {
				
			}
		}
	}
}

// @lc code=end
