/*
 * @lc app=leetcode id=275 lang=golang
 *
 * [275] H-Index II
 */

package golang

import "testing"

func Test_hIndexII(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[]int{0, 1, 3, 5, 6}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndexII(tt.args.citations); got != tt.want {
				t.Errorf("hIndexII() = %v, want %v", got, tt.want)
			}
		})
	}
}
