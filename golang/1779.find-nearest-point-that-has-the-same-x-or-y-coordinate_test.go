/*
 * @lc app=leetcode id=1779 lang=golang
 *
 * [1779] Find Nearest Point That Has the Same X or Y Coordinate
 */

package golang

import "testing"

func Test_nearestValidPoint(t *testing.T) {
	type args struct {
		x      int
		y      int
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestValidPoint(tt.args.x, tt.args.y, tt.args.points); got != tt.want {
				t.Errorf("nearestValidPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
