/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (37.72%)
 * Likes:    3803
 * Dislikes: 190
 * Total Accepted:    792.8K
 * Total Submissions: 2.1M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output:
 want:true,

 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output:
 want:true,

 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output:
 want:false,

 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output:
 want:false,

 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output:
 want:true,

 *
 *
*/

package leetcode

import (
	"testing"
)

func Test_isValid(t *testing.T) {
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
			args: args{"()"},
			want: true,
		},
		{
			name: "Test 2",
			args: args{"()[]{}"},
			want: true,
		},
		{
			name: "Test 3",
			args: args{"(]"},
			want: false,
		},
		{
			name: "Test 4",
			args: args{"([)]"},
			want: false,
		},
		{
			name: "Test 5",
			args: args{"{[]}"},
			want: true,
		},
		{
			name: "Test 6",
			args: args{"["},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
