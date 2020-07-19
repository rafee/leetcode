/*
 * @lc app=leetcode id=282 lang=golang
 *
 * [282] Expression Add Operators
 */

package golang

import (
	"strconv"
)

// @lc code=start
type Operation int

const (
	Add Operation = iota
	Sub
	Mul
	Last
)

func (o Operation) String() string {
	return [...]string{"+", "-", "*"}[o]
}

func addOperators(num string, target int) []string {
	if len(num) == 1 {
		val, _ := strconv.Atoi(num)
		if val == target {
			return []string{num}
		}
	}
	result := make([]string, 0)
	for i := Add; i < Last; i++ {
		// result = append(result, getNextExpr())
	}
	return result
}

// @lc code=end
