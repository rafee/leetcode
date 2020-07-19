/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */

package golang

// @lc code=start
func addStrings(num1 string, num2 string) string {
	var carry byte
	N1, N2 := len(num1), len(num2)
	if N2 < N1 {
		return addStrings(num2, num1)
	}
	result := make([]byte, N2+1)
	for i := 0; i < N1; i++ {
		digit := num1[N1-i-1] + num2[N2-i-1] - '0' + carry
		digit, carry = checkCarry(digit)
		result[N2-i] = digit
	}

	for i := N1; i < N2; i++ {
		digit := num2[N2-i-1] + carry
		digit, carry = checkCarry(digit)
		result[N2-i] = digit
	}

	if carry == 1 {
		result[0] = '1'
	} else {
		result = result[1:]
	}
	return string(result)
}

func checkCarry(digit byte) (byte, byte) {
	var carry byte
	if digit > '9' {
		digit -= 10
		carry = 1
	} else {
		carry = 0
	}
	return digit, carry
}

// @lc code=end
