/*
 * @lc app=leetcode id=540 lang=golang
 *
 * [540] Single Element in a Sorted Array
 */

package golang

import "testing"

func Test_singleNonDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{[]int{3, 3, 7, 7, 10, 11, 11}},
			want: 10,
		},
		{
			name: "Test 3",
			args: args{[]int{1, 1, 2}},
			want: 2,
		},
		{
			name: "Test 4",
			args: args{[]int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNonDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
