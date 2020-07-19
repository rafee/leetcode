/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

package golang

import "strconv"

// @lc code=start
func numDecodings(s string) int {
	return numDecodingsIterative(s)
}

func numDecodingsIterative(s string) int {
	result := make([]int, len(s)+1)
	result[len(s)] = 1
	if s[len(s)-1] != '0' {
		result[len(s)-1] = 1
	} else {
		result[len(s)-1] = 0
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}
		result[i] = result[i+1]
		num, _ := strconv.Atoi(s[i : i+2])
		if num <= 26 {
			result[i] += result[i+2]
		}
	}
	return result[0]
}

func numDecodingsRecursive(s string) int {
	if s == "" {
		return 1
	} else if s[0] == '0' {
		return 0
	} else if len(s) == 1 {
		return 1
	}
	result := numDecodingsRecursive(s[1:])
	num, _ := strconv.Atoi(s[:2])
	if num <= 26 {
		result += numDecodingsRecursive(s[2:])
	}
	return result
}

// @lc code=end
