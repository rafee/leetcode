/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 *
 * https://leetcode.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (44.57%)
 * Likes:    1426
 * Dislikes: 61
 * Total Accepted:    259K
 * Total Submissions: 566.4K
 * Testcase Example:  '[1,2,2]'
 *
 * Given a collection of integers that might contain duplicates, nums, return
 * all possible subsets (the power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: [1,2,2]
 * Output:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
 *
 */

package golang

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{[]int{1, 2, 2}},
			want: [][]int{{}, {2}, {2, 2}, {1}, {1, 2}, {1, 2, 2}},
		},
		{
			name: "Test 2",
			args: args{[]int{9, 0, 3, 5, 7}},
			want: [][]int{{}, {9}, {7}, {7, 9}, {5}, {5, 9}, {5, 7}, {5, 7, 9}, {3}, {3, 9}, {3, 7}, {3, 7, 9}, {3, 5}, {3, 5, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0}, {0, 9}, {0, 7}, {0, 7, 9}, {0, 5}, {0, 5, 9}, {0, 5, 7}, {0, 5, 7, 9}, {0, 3}, {0, 3, 9}, {0, 3, 7}, {0, 3, 7, 9}, {0, 3, 5}, {0, 3, 5, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsWithDup(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				assert.ElementsMatch(t, got, tt.want)
			}
		})
	}
}

func Test_doSubsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
		cur  []int
		res  *[][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doSubsetsWithDup(tt.args.nums, tt.args.cur, tt.args.res)
		})
	}
}
