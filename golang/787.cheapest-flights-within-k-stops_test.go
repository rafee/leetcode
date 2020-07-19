/*
 * @lc app=leetcode id=787 lang=golang
 *
 * [787] Cheapest Flights Within K Stops
 */

package golang

import "testing"

func Test_findCheapestPrice(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		K       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1},
			want: 200,
		},
		{
			name: "Test 2",
			args: args{3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCheapestPrice(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.K); got != tt.want {
				t.Errorf("findCheapestPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
