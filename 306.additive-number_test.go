/*
 * @lc app=leetcode id=306 lang=golang
 *
 * [306] Additive Number
 */

package leetcode

import (
	"testing"
)

func Test_isAdditiveNumber(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{"112358"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdditiveNumber(tt.args.num); got != tt.want {
				t.Errorf("isAdditiveNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAdditive(t *testing.T) {
	type args struct {
		num1 string
		num2 string
		next string
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
			if got := isAdditive(tt.args.num1, tt.args.num2, tt.args.next); got != tt.want {
				t.Errorf("isAdditive() = %v, want %v", got, tt.want)
			}
		})
	}
}
