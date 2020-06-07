/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (55.12%)
 * Likes:    2826
 * Dislikes: 212
 * Total Accepted:    577.9K
 * Total Submissions: 1M
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */

package leetcode

// @lc code=start
func majorityElement(nums []int) int {
	numMap := make(map[int]int)
	maxCount := len(nums) / 2
	for _, num := range nums {
		numMap[num]++
		if numMap[num] > maxCount {
			return num
		}
	}
	return -1
}

// @lc code=end
