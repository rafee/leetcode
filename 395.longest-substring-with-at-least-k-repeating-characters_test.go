/*
 * @lc app=leetcode id=395 lang=golang
 *
 * [395] Longest Substring with At Least K Repeating Characters
 */

package golang

import "testing"

func Test_longestSubstring(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{"aaabb", 3},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{"ababbc", 2},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("longestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
