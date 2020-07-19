/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{[]int{2, 0, 2, 1, 1, 0}},
			want: []int{0, 0, 1, 1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !reflect.DeepEqual(tt.want, tt.args.nums) {
				t.Error("Wanted ", tt.want, "\nGot ", tt.args.nums)
			}
		})
	}
}
