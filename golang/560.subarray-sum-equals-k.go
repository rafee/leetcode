/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 *
 * https://leetcode.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (43.22%)
 * Likes:    3818
 * Dislikes: 123
 * Total Accepted:    240.1K
 * Total Submissions: 548.6K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * Given an array of integers and an integer k, you need to find the total
 * number of continuous subarrays whose sum equals to k.
 *
 * Example 1:
 *
 * Input:nums = [1,1,1], k = 2
 * Output: 2
 *
 *
 *
 * Note:
 *
 * The length of the array is in range [1, 20,000].
 * The range of numbers in the array is [-1000, 1000] and the range of the
 * integer k is [-1e7, 1e7].
 *
 *
 *
 */

package golang

// @lc code=start
func subarraySum(nums []int, k int) int {
	sumMap := make(map[int]int)
	sumMap[0] = 1
	count, sum := 0, 0
	for _, num := range nums {
		sum += num
		count += sumMap[sum-k]
		sumMap[sum]++
	}
	return count
}

// @lc code=end
