/*
 * @lc app=leetcode id=886 lang=golang
 *
 * [886] Possible Bipartition
 */

package golang

import "testing"

func Test_possibleBipartition(t *testing.T) {
	type args struct {
		N        int
		dislikes [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartition(tt.args.N, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_visitNext(t *testing.T) {
	type args struct {
		graph  [][]bool
		colors map[int]bool
		src    int
		color  bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := visitNext(tt.args.graph, tt.args.colors, tt.args.src, tt.args.color); got != tt.want {
				t.Errorf("visitNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
