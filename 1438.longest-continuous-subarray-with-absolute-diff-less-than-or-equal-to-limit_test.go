/*
 * @lc app=leetcode id=1438 lang=golang
 *
 * [1438] Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
 */

package leetcode

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[]int{8, 2, 4, 7}, 4},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{[]int{10, 1, 2, 4, 7, 2}, 5},
			want: 4,
		},
		{
			name: "Test 3",
			args: args{[]int{4, 2, 2, 2, 4, 4, 2, 2}, 0},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
