/*
 * @lc app=leetcode id=918 lang=golang
 *
 * [918] Maximum Sum Circular Subarray
 */

package leetcode

import "testing"

func Test_maxSubarraySumCircular(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[]int{1, -2, 3, -2}},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{[]int{5,-3,5}},
			want: 10,
		},
		{
			name: "Test 3",
			args: args{[]int{3,-1,2,-1}},
			want: 4,
		},
		{
			name: "Test 4",
			args: args{[]int{3,-2,2,-3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircular(tt.args.A); got != tt.want {
				t.Errorf("maxSubarraySumCircular() = %v, want %v", got, tt.want)
			}
		})
	}
}
