/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

package golang

// @lc code=start
func isAnagram(s string, t string) bool {
	// sMap := make(map[rune]int)
	// for _, r := range s {
	// 	sMap[r]++
	// }
	// for _, r := range t {
	// 	sMap[r]--
	// 	if sMap[r] < 0 {
	// 		return false
	// 	} else if sMap[r] == 0 {
	// 		delete(sMap, r)
	// 	}
	// }
	// return len(sMap) == 0
	if len(s) != len(t) {
		return false
	}
	count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		count[int(s[i])-int('a')]++
		count[int(t[i])-int('a')]--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
