/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 */

package golang

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// sort.Ints(nums2)
	numMap := make(map[int]int)
	stack := make([]int, 0)
	for _, num := range nums2 {
		for top := len(stack) - 1; len(stack) != 0 && stack[top] < num; top =
			len(stack) - 1 {
			numMap[stack[top]] = num
			stack = stack[:top]
		}
		stack = append(stack, num)
	}
	result := make([]int, len(nums1))
	for i, num := range nums1 {
		res, found := numMap[num]
		if !found {
			res = -1
		}
		result[i] = res
	}
	return result
}

// @lc code=end
