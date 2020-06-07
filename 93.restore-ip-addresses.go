/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

package leetcode

import (
	"strconv"
	"strings"
)

// @lc code=start
func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	if len(s) < 4 || len(s) > 12 {
		return result
	}
	segments := make([]string, 0)
	iPAdrressHelper(s, segments, &result)
	return result
}

func iPAdrressHelper(s string, segments []string, result *[]string) {
	if len(segments) == 3 {
		if chkNum(s) {
			return
		}
		segments = append(segments, s)
		IP := strings.Join(segments, ".")
		*result = append(*result, IP)
	} else {
		for i, cur := 0, int(s[0]-'0'); ; {
			iPAdrressHelper(s[i+1:], append(segments, s[:i+1]), result)
			i++
			// len(s)-3+len(segments) checks whether we have enough characters
			// in string to go deeper
			if i == 3 || i == (len(s)-3+len(segments)) {
				return
			}
			cur *= 10
			cur += int(s[i] - '0')
			if s[0] == '0' || cur > 255 {
				return
			}
		}
	}
}

// chkNum checks for leading zeros or whether the number is greater than 255.
// Returns true for ineligible numbers
func chkNum(s string) bool {
	num, _ := strconv.Atoi(s)
	return num > 255 || (s[0] == '0' && len(s) > 1)
}

// @lc code=end
