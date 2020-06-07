/*
 * @lc app=leetcode id=306 lang=golang
 *
 * [306] Additive Number
 */

package leetcode

// @lc code=start
func isAdditiveNumber(num string) bool {
	for i := 0; i < len(num)-2; i++ {
		for j := i + 2; j < len(num); j++ {
			if isAdditive(num[:i+1], num[i+1:j], num[j:]) {
				return true
			}
		}
	}
	return false
}

func isAdditive(num1 string, num2 string, next string) bool {
	if num2[0] == '0' && len(num2) > 1 {
		return false
	}
	curNum := addStrings(num1, num2)
	if len(curNum) > len(next) {
		return false
	}
	str := next[:len(curNum)]
	var nextStr string
	if curNum == str {
		if len(curNum) == len(next) {
			return true
		}
		nextStr = next[len(curNum):]
	}
	return isAdditive(num2, curNum, nextStr)
}
// @lc code=end
