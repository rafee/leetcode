/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

package leetcode

import (
	"strconv"
)

// @lc code=start
func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	bArr := make([]byte, len(s)+3)
	iPAdrressHelper(s, 0, bArr, &result)
	return result
}

func iPAdrressHelper(s string, steps int, bArr []byte, result *[]string) {
	if steps == 3 {
		num, _ := strconv.Atoi(s)
		if num > 255 || strconv.Itoa(num) != s {
			return
		}

		// *result = append(*result, sb.String())
	} else {
		cur := int(s[0] - '0')
		steps++
		for i := 1; i < 3 && i < len(s)-1; i++ {
			// sb.WriteByte(s[i])
			// sb.WriteByte('.')
			// iPAdrressHelper(s[i+1:], steps, sb, result)
			// tmp := sb.String()
			// tmp = tmp[:len(tmp)-1]
			// sb.Reset()
			// sb.WriteString(tmp)
			cur *= 10
			cur += int(s[i] - '0')
			if cur > 255 || cur == 0 {
				return
			}
		}
	}
	return
}

// @lc code=end
