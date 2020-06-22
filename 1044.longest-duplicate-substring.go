/*
 * @lc app=leetcode id=1044 lang=golang
 *
 * [1044] Longest Duplicate Substring
 */

package leetcode

// @lc code=start
func longestDupSubstring(str string) string {
	left, right := 0, len(str)
	greedstr := ""
	for left < right {
		middle := (left + right) / 2
		d := findDuplicate(str, middle)

		// search to the left
		if d == "" {
			if right == middle {
				break
			}
			right = middle
		} else {

			// search to the right
			if left == middle {
				break
			}
			left = middle
			greedstr = d
		}
	}
	return greedstr
}

func findDuplicate(str string, length int) string {
	dic := make(map[string]bool, len(str)-length)
	for i := 0; i <= len(str)-length; i++ {
		s := str[i : i+length]
		if dic[s] {
			return s
		}
		dic[s] = true
	}
	return ""
}

// @lc code=end
