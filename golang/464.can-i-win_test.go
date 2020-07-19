/*
 * @lc app=leetcode id=464 lang=golang
 *
 * [464] Can I Win
 */

package golang

import (
	"testing"
)

func Test_canIWin(t *testing.T) {
	type args struct {
		maxChoosableInteger int
		desiredTotal        int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{10, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canIWin(tt.args.maxChoosableInteger, tt.args.desiredTotal); got != tt.want {
				t.Errorf("canIWin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatBoolArrToInt(t *testing.T) {
	type args struct {
		arr []bool
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatBoolArrToInt(tt.args.arr); got != tt.want {
				t.Errorf("formatBoolArrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkWin(t *testing.T) {
	type args struct {
		desiredTotal int
		used         []bool
		state        map[int32]bool
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
			if got := checkWin(tt.args.desiredTotal, tt.args.used, tt.args.state); got != tt.want {
				t.Errorf("checkWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
