/*
 * @lc app=leetcode id=464 lang=golang
 *
 * [464] Can I Win
 */

package leetcode

// @lc code=start
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	sum := (maxChoosableInteger + 1) / 2 * maxChoosableInteger
	if sum < desiredTotal {
		return false
	}
	if desiredTotal <= 0 {
		return true
	}

	var used = make([]bool, maxChoosableInteger+1)
	var state = map[int32]bool{}
	return checkWin(desiredTotal, used, state)
}

func checkWin(desiredTotal int, used []bool, state map[int32]bool) bool {
	if desiredTotal <= 0 {
		return false
	}
	val := formatBoolArrToInt(used)
	_, found := state[val]
	if !found {
		for i := 1; i < len(used); i++ {
			if !used[i] {
				used[i] = true
				if !checkWin(desiredTotal-i, used, state) {
					state[val] = true
					used[i] = false
					return true
				}
				used[i] = false
			}
		}
		state[val] = false
	}
	return state[val]
}

func formatBoolArrToInt(arr []bool) int32 {
	var num int32 = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] {
			num |= (1 << i)
		}
	}
	return num
}

// @lc code=end
