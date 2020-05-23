/*
 * @lc app=leetcode id=329 lang=golang
 *
 * [329] Longest Increasing Path in a Matrix
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_longestIncreasingPath(t *testing.T) {
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
			args: args{[][]int{{1, 2}}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath(tt.args.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextValid(t *testing.T) {
	type args struct {
		i int
		j int
		X int
		Y int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextValid(tt.args.i, tt.args.j, tt.args.X, tt.args.Y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkLongSeq(t *testing.T) {
	type args struct {
		matrix    [][]int
		seqMatrix [][]int
		i         int
		j         int
		X         int
		Y         int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkLongSeq(tt.args.matrix, tt.args.seqMatrix, tt.args.i, tt.args.j, tt.args.X, tt.args.Y)
		})
	}
}
