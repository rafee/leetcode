/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

package golang

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{"12"},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{"226"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numDecodingsIterative(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodingsIterative(tt.args.s); got != tt.want {
				t.Errorf("numDecodingsIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numDecodingsRecursive(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodingsRecursive(tt.args.s); got != tt.want {
				t.Errorf("numDecodingsRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
