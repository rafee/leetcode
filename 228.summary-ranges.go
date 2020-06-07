/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

package leetcode

import (
	"fmt"
)

// @lc code=start
func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	if len(nums) == 0 {
		return result
	}
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if (nums[i] - nums[i-1]) > 1 {
			str := getRngStr(nums[i-1], pre)
			pre = nums[i]
			result = append(result, str)
		}
	}
	str := getRngStr(nums[len(nums)-1], pre)
	result = append(result, str)
	return result
}

func getRngStr(num int, pre int) string {
	var str string
	if pre == num {
		str = fmt.Sprintf("%d", pre)
	} else {
		str = fmt.Sprintf("%d->%d", pre, num)
	}
	return str
}

// @lc code=end
