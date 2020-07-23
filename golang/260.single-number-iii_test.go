/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 */

package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumberIII(t *testing.T) {
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
			args: args{[]int{1, 2, 1, 3, 2, 5}},
			want: []int{3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, singleNumberIII(tt.args.nums))
		})
	}
}
