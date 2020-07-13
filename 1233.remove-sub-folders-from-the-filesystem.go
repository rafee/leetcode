/*
 * @lc app=leetcode id=1233 lang=golang
 *
 * [1233] Remove Sub-Folders from the Filesystem
 */

package leetcode

import (
	"sort"
	"strings"
)

// @lc code=start
func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	strMap := make(map[string]bool)
	result := make([]string, 0)
	for _, path := range folder {
		subs := strings.Split(path, "/")
		subs = subs[1:]
		var sb strings.Builder
		for _, sub := range subs {
			sb.WriteString("/")
			sb.WriteString(sub)
			if strMap[sb.String()] {
				break
			}
		}
		if sb.String() == path {
			result = append(result, path)
			strMap[path] = true
		}
	}
	return result
}

// @lc code=end
