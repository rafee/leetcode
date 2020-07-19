/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (57.70%)
 * Likes:    2936
 * Dislikes: 70
 * Total Accepted:    488.8K
 * Total Submissions: 842.2K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct integers, nums, return all possible subsets (the
 * power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: nums = [1,2,3]
 * Output:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
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
			args: args{[]int{}},
			want: [][]int{{}},
		},
		{
			name: "Test 2",
			args: args{[]int{1}},
			want: [][]int{{}, {1}},
		},
		{
			name: "Test 3",
			args: args{[]int{1, 2, 3}},
			want: [][]int{
				{},
				{1},
				{2},
				{3},
				{2, 3},
				{1, 2},
				{1, 3},
				{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
