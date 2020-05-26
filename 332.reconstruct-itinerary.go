/*
 * @lc app=leetcode id=332 lang=golang
 *
 * [332] Reconstruct Itinerary
 */

package leetcode

import "sort"

// @lc code=start
func findItinerary(tickets [][]string) []string {
	destMap := make(map[string][]string)
	for _, ticket := range tickets {
		_, found := destMap[ticket[0]]
		if found {
			destMap[ticket[0]] = append(destMap[ticket[0]], ticket[1])
		} else {
			destMap[ticket[0]] = []string{ticket[1]}
		}
	}

	result := []string{"JFK"}
	for dest := "JFK"; destMap[dest] != nil; {
		stringSlice := destMap[dest]
		sort.Strings(stringSlice)
		pre := dest
		dest = stringSlice[0]
		if len(stringSlice) == 1 {
			destMap[pre] = nil
		} else {
			destMap[pre] = stringSlice[1:]
		}
		result = append(result, dest)
	}
	return result
}

// @lc code=end
