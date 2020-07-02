/*
 * @lc app=leetcode id=335 lang=golang
 *
 * [335] Self Crossing
 */

package leetcode

import "testing"

func Test_isSelfCrossing(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{[]int{2, 1, 1, 2}},
			want: true,
		},
		{
			name: "Test 2",
			args: args{[]int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "Test 3",
			args: args{[]int{1, 1, 1, 1}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSelfCrossing(tt.args.x); got != tt.want {
				t.Errorf("isSelfCrossing() = %v, want %v", got, tt.want)
			}
		})
	}
}
