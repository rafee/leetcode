/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

package golang

import (
	"testing"
)

func Test_isAlphaNum(t *testing.T) {
	type args struct {
		b byte
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
			if got := isAlphaNum(tt.args.b); got != tt.want {
				t.Errorf("isAlphaNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{"race a car"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeString(tt.args.s); got != tt.want {
				t.Errorf("isPalindromeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
