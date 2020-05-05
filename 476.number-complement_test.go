/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 *
 * https://leetcode.com/problems/number-complement/description/
 *
 * algorithms
 * Easy (62.97%)
 * Likes:    729
 * Dislikes: 77
 * Total Accepted:    137.6K
 * Total Submissions: 215.7K
 * Testcase Example:  '5'
 *
 * Given a positive integer, output its complement number. The complement
 * strategy is to flip the bits of its binary representation.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 5
 * Output: 2
 * Explanation: The binary representation of 5 is 101 (no leading zero bits),
 * and its complement is 010. So you need to output 2.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: 1
 * Output: 0
 * Explanation: The binary representation of 1 is 1 (no leading zero bits), and
 * its complement is 0. So you need to output 0.
 *
 *
 *
 *
 * Note:
 *
 *
 * The given integer is guaranteed to fit within the range of a 32-bit signed
 * integer.
 * You could assume no leading zero bit in the integer’s binary
 * representation.
 * This question is the same as 1009:
 * https://leetcode.com/problems/complement-of-base-10-integer/
 *
 *
 */

package leetcode

import "testing"

func Test_findComplement(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{5},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findComplement(tt.args.num); got != tt.want {
				t.Errorf("findComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
