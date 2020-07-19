/*
 * @lc app=leetcode id=332 lang=golang
 *
 * [332] Reconstruct Itinerary
 */

package golang

import (
	"fmt"
	"sort"
)

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

	for _, destinations := range destMap {
		sort.Strings(destinations)
		// fmt.Println(src,destinations)
	}

	result := visitAirport("JFK", destMap, []string{}, len(tickets))
	return result
}

func visitAirport(src string, destMap map[string][]string, pre []string,
	ttl int) []string {
	curRes := append(pre, src)
	destinations := make([]string, len(destMap[src]))
	fmt.Println(curRes, ttl, destinations)
	copy(destinations, destMap[src])
	if len(destinations) == 0 {
		if ttl != 0 {
			return nil
		}
		return curRes
	}

	for i, dest := range destinations {
		tmp := make([]string, len(destinations)-1)
		copy(tmp[:i], destinations[:i])
		copy(tmp[i:], destinations[i+1:])
		destMap[src] = tmp
		result := visitAirport(dest, destMap, curRes, ttl-1)
		if result != nil {
			return result
		}
	}
	return nil
}

// @lc code=end
