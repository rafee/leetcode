/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

package golang

import (
	"strconv"
	"strings"
)

// @lc code=start
func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		var sb strings.Builder
		if i%3 == 0 {
			sb.WriteString("Fizz")
		}
		if i%5 == 0 {
			sb.WriteString("Buzz")
		}
		if i%3 != 0 && i%5 != 0 {
			sb.WriteString(strconv.Itoa(i))
		}
		result[i-1] = sb.String()
	}
	return result
}

// @lc code=end
