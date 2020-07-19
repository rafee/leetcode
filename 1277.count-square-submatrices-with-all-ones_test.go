/*
 * @lc app=leetcode id=1277 lang=golang
 *
 * [1277] Count Square Submatrices with All Ones
 */

package golang

import "testing"

func Test_countSquares(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSquares(tt.args.matrix); got != tt.want {
				t.Errorf("countSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
