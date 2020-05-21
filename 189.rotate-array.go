/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

package leetcode

// @lc code=start
func rotate(nums []int, k int) {
	numLen := len(nums)
	for count, i := 0, 0; count < numLen; {
		j := (i + k) % numLen
		num := nums[i]
		for i != j {
			tmp := nums[j]
			nums[j] = num
			j = (j + k) % numLen
			num = tmp
			count++
		}
		i++
	}
}

// @lc code=end
