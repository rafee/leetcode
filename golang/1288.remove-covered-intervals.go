/*
 * @lc app=leetcode id=1288 lang=golang
 *
 * [1288] Remove Covered Intervals
 */

package golang

import (
	"fmt"
	"sort"
)

// @lc code=start
func removeCoveredIntervals(intervals [][]int) int {
	sort.Sort(intervalsType(intervals))
	fmt.Println(intervals)
	count:=0
	
	high:=0
	for _, i := range intervals {
		if i[1]>high{
			high=i[1]
			count++
		}
	}
	return count
}

type intervalsType [][]int

func (x intervalsType) Len() int{
	return len(x)
}

func (x intervalsType) Swap(i,j int){
	x[i],x[j]=x[j],x[i]
}

func (x intervalsType) Less(i,j int)bool{
	if x[i][0]!=x[j][0]{
		return x[i][0]<x[j][0]
	}
	return x[i][1]>x[j][1]
}
// @lc code=end

