/*
 * @lc app=leetcode id=1790 lang=golang
 *
 * [1790] Check if One String Swap Can Make Strings Equal
 */

package golang

// @lc code=start
func areAlmostEqual(s1 string, s2 string) bool {
	var c1, c2 byte
	mismatch := false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if c1 == 0 {
				c1, c2 = s1[i], s2[i]
			} else {
				if mismatch || s1[i] != c2 || s2[i] != c1 {
					return false
				}
				mismatch = true
				c1 = 'A'
			}
		}
	}
	if c1 == 'A' || c1 == 0 {
		return true
	}
	return false
}

// @lc code=end
