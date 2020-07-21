/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */

package golang

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{[]int{10, 2}},
			want: "210",
		},
		{
			name: "Test 2",
			args: args{[]int{3, 30, 34, 5, 9}},
			want: "9534330",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
