/*
 * @lc app=leetcode id=1044 lang=golang
 *
 * [1044] Longest Duplicate Substring
 */

package golang

import (
	"testing"
)

func Test_longestDupSubstring(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{"banana"},
			want: "ana",
		},
		{
			name: "Test 2",
			args: args{"abcd"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDupSubstring(tt.args.S); got != tt.want {
				t.Errorf("longestDupSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicate(t *testing.T) {
	type args struct {
		str    string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.str, tt.args.length); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
