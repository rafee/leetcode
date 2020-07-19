/*
 * @lc app=leetcode id=914 lang=golang
 *
 * [914] X of a Kind in a Deck of Cards
 */

package golang

import "testing"

func Test_hasGroupsSizeX(t *testing.T) {
	type args struct {
		deck []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{[]int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasGroupsSizeX(tt.args.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX() = %v, want %v", got, tt.want)
			}
		})
	}
}
