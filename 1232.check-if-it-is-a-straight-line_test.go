/*
 * @lc app=leetcode id=1232 lang=golang
 *
 * [1232] Check If It Is a Straight Line
 */

package leetcode

import "testing"

func Test_checkStraightLine(t *testing.T) {
	type args struct {
		coordinates [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}},
			want: true,
		},
		{
			name: "Test 2",
			args: args{[][]int{{1, 2}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}},
			want: false,
		},
		{
			name: "Test 3",
			args: args{[][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStraightLine(tt.args.coordinates); got != tt.want {
				t.Errorf("checkStraightLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
