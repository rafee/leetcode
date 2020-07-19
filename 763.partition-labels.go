/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 */

package golang

// @lc code=start
func partitionLabels(S string) []int {
	charMap := make([]int, 26)
	for i := 0; i < len(S); i++ {
		charMap[S[i]-'a'] = i
	}

	res := make([]int, 0)
	for i, mx, pre := 0, 0, 0; i < len(S); i++ {
		if charMap[S[i]-'a'] > mx {
			mx = charMap[S[i]-'a']
		}
		if (i + 1) > mx {
			res = append(res, mx+1-pre)
			pre = mx + 1
		}
	}

	return res
}

// @lc code=end
