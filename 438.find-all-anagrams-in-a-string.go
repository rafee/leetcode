/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

package leetcode

// @lc code=start
func findAnagrams(s string, p string) []int {
	lenS, lenP := len(s), len(p)
	if lenS < lenP {
		return nil
	}
	var chk, run [26]int
	for i := 0; i < lenP; i++ {
		chk[p[i]-'a']++
		run[s[i]-'a']++
	}
	result := make([]int, 0)
	for i := lenP; ; i++ {
		if chk == run {
			result = append(result, i-lenP)
		}
		if i == lenS {
			break
		}
		run[s[i-lenP]-'a']--
		run[s[i]-'a']++
	}
	return result
}

// @lc code=end
