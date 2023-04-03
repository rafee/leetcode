/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

package golang

// @lc code=start
func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	lenS, lenT := len(s), len(t)
	if lenS != lenT {
		return false
	}
	for i := 0; i < lenS; i++ {
		r1, found1 := m1[s[i]]
		r2, found2 := m2[t[i]]
		if found1 != found2 {
			return false
		}
		if found1 {
			if r1 != t[i] || r2 != s[i] {
				return false
			}
		} else {
			m1[s[i]] = t[i]
			m2[t[i]] = s[i]
		}
	}
	return true
}

// @lc code=end
