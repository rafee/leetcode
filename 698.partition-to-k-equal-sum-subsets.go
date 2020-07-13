/*
 * @lc app=leetcode id=698 lang=golang
 *
 * [698] Partition to K Equal Sum Subsets
 */

package leetcode

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	s, max := 0, 0
	for _, num := range nums {
		s += num
		if num > max {
			max = num
		}
	}
	if s%k != 0 || max > s/k {
		return false
	}

	for i, target := 0, s/k; i < len(nums); i++ {
		target -= nums[i]
		if target == 0 {
			target = s / k
		}
	}
	return true
}

// @lc code=end
