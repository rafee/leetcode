/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

package leetcode

// @lc code=start
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		return addBinary(b, a)
	}

	result := make([]byte, len(a)+1)
	carry := 0
	for i, j := len(a)-1, len(b)-1; i >= 0; i, j = i-1, j-1 {
		if j >= 0 {
			digit := int(a[i]-'0') + int(b[j]-'0') + carry
			carry = digit / 2
			digit = digit % 2
			result[i+1] = byte(digit) + '0'
		} else {
			digit := int(a[i]-'0') + carry
			carry = digit / 2
			digit = digit % 2
			result[i+1] = byte(digit) + '0'
		}
	}
	if carry == 1 {
		result[0] = '1'
		return string(result)
	}
	return string(result[1:])
}

// @lc code=end
