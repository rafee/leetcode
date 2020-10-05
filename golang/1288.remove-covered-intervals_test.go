/*
 * @lc app=leetcode id=1288 lang=golang
 *
 * [1288] Remove Covered Intervals
 */

package golang

import "testing"

func Test_removeCoveredIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[][]int{{1,4},{3,6},{2,8}}},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{[][]int{{1,4},{2,3}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeCoveredIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("removeCoveredIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intervalsType_Len(t *testing.T) {
	tests := []struct {
		name string
		x    intervalsType
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.Len(); got != tt.want {
				t.Errorf("intervalsType.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intervalsType_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		x    intervalsType
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.x.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_intervalsType_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		x    intervalsType
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.x.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("intervalsType.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
