/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 *
 * https://leetcode.com/problems/ugly-number/description/
 *
 * algorithms
 * Easy (41.23%)
 * Likes:    389
 * Dislikes: 572
 * Total Accepted:    193.3K
 * Total Submissions: 468.7K
 * Testcase Example:  '6'
 *
 * Write a program to check whether a given number is an ugly number.
 *
 * Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
 *
 * Example 1:
 *
 *
 * Input: 6
 * Output: true
 * Explanation: 6 = 2 × 3
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: true
 * Explanation: 8 = 2 × 2 × 2
 *
 *
 * Example 3:
 *
 *
 * Input: 14
 * Output: false
 * Explanation: 14 is not ugly since it includes another prime factor 7.
 *
 *
 * Note:
 *
 *
 * 1 is typically treated as an ugly number.
 * Input is within the 32-bit signed integer range: [−2^31,  2^31 − 1].
 *
 */

package golang

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{6},
			want: true,
		},
		{
			name: "Test 2",
			args: args{8},
			want: true,
		},
		{
			name: "Test 3",
			args: args{14},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.num); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
