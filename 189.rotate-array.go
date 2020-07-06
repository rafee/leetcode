/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

package leetcode

// @lc code=start
func rotate(nums []int, k int) {
	k = k % len(nums)
	for start, count := 0, 0; count < len(nums); {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % len(nums)
			nums[next], prev = prev, nums[next]
			current = next
			count++
			if start == current {
				break
			}
		}
	}
}

// @lc code=end
