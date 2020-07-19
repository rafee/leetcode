/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{[]int{4, 1, 2}, []int{1, 3, 4, 2}},
			want: []int{-1, 3, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
