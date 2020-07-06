/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

package leetcode

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)

		})
	}
}
