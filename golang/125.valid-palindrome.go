/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

package golang

import "strings"

// @lc code=start
// isPalindromeString checks whether a string is palindrome considering
// alphanumeric characters only while ignoring case
func isPalindromeString(s string) bool {
	s = strings.ToLower(s)
	for left, right := 0, len(s)-1; left < right; {
		if !isAlphaNum(s[left]) {
			left++
			continue
		}
		if !isAlphaNum(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--

	}
	return true
}

// isAlphaNum checks whether the byte is a lowercase character or number
func isAlphaNum(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

// @lc code=end
