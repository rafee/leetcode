/*
 * @lc app=leetcode id=1344 lang=golang
 *
 * [1344] Angle Between Hands of a Clock
 */

package leetcode

import "testing"

func Test_angleClock(t *testing.T) {
	type args struct {
		hour    int
		minutes int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{12, 30},
			want: 165,
		},
		{
			name: "Test 2",
			args: args{3, 30},
			want: 75,
		},
		{
			name: "Test 3",
			args: args{3, 15},
			want: 7.5,
		},
		{
			name: "Test 4",
			args: args{1, 57},
			want: 76.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := angleClock(tt.args.hour, tt.args.minutes); got != tt.want {
				t.Errorf("angleClock() = %v, want %v", got, tt.want)
			}
		})
	}
}
