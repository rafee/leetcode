/*
 * @lc app=leetcode id=335 lang=golang
 *
 * [335] Self Crossing
 */

package golang

import (
	"testing"
)

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

func Test_checkCross(t *testing.T) {
	type args struct {
		pos1 [2]int
		pos2 [2]int
		pos3 [2]int
		pos4 [2]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkCross(tt.args.pos1, tt.args.pos2, tt.args.pos3, tt.args.pos4); got != tt.want {
				t.Errorf("checkCross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sign(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sign(tt.args.n); got != tt.want {
				t.Errorf("sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
