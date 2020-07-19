/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

package golang

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	// for i := m + n - 1; i >= n; i-- {
	// 	nums1[i] = nums1[i-n]
	// }

	// for i, j := n, 0; j < n; {
	// 	if i < (m+n) && nums1[i] < nums2[j] {
	// 		nums1[i+j-n] = nums1[i]
	// 		i++
	// 	} else {
	// 		nums1[i+j-n] = nums2[j]
	// 		j++
	// 	}
	// }

	for i, j := m-1, n-1; j >= 0; {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}
}

// @lc code=end
