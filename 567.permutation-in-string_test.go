/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

package leetcode

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{"ab", "eidbaooo"},
			want: true,
		},
		{
			name: "Test 2",
			args: args{"ab", "eidboaooo"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
