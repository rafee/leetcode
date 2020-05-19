/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

package leetcode

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	lenS1, lenS2 := len(s1), len(s2)
	if lenS1 > lenS2 {
		return false
	}
	var chk, run [26]int
	for i := 0; i < lenS1; i++ {
		chk[s1[i]-'a']++
		run[s2[i]-'a']++
	}
	for i := lenS1; ; i++ {
		if chk == run {
			return true
		}
		if i == lenS2 {
			break
		}
		run[s2[i-lenS1]-'a']--
		run[s2[i]-'a']++
	}
	return false
}

// @lc code=end
