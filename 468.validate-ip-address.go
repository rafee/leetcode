/*
 * @lc app=leetcode id=468 lang=golang
 *
 * [468] Validate IP Address
 */

package leetcode

import (
	"strings"
)

// @lc code=start
func validIPAddress(IP string) string {
	if segs := strings.Split(IP, "."); len(segs) == 4 {
		return checkIPv4(segs)
	} else if segs := strings.Split(IP, ":"); len(segs) == 8 {
		return checkIPv6(segs)
	} else {
		return "Neither"
	}
}

func checkIPv4(segments []string) string {
	for _, seg := range segments {
		if len(seg) == 0 {
			return "Neither"
		}
		checkRange4 := func(b int) bool {
			return b < 10
		}
		num := int(seg[0] - '0')
		if !checkRange4(num) {
			return "Neither"
		}
		for i := 1; i < len(seg); i++ {
			num *= 10
			b := int(seg[i] - '0')
			if !checkRange4(b) {
				return "Neither"
			}
			num += b
			if num == b || num > 255 {
				return "Neither"
			}
		}
	}
	return "IPv4"
}

func checkIPv6(segments []string) string {
	for _, seg := range segments {
		if len(seg) == 0 || len(seg) > 4 {
			return "Neither"
		}
		checkRange6 := func(b byte) bool {
			return (b >= '0' && b <= '9') || (b >= 'A' && b <= 'F') || (b >= 'a' && b <= 'f')
		}
		for i := 0; i < len(seg); i++ {
			if !checkRange6(seg[i]) {
				return "Neither"
			}
		}
	}
	return "IPv6"
}

// @lc code=end
