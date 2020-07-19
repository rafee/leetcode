/*
 * @lc app=leetcode id=406 lang=golang
 *
 * [406] Queue Reconstruction by Height
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	type args struct {
		people [][]int
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
			if got := reconstructQueue(tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	type args struct {
		sli   *[][]int
		elem  []int
		index int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insert(tt.args.sli, tt.args.elem, tt.args.index)
		})
	}
}

func TestByHeightAndRank_Len(t *testing.T) {
	tests := []struct {
		name string
		b    ByHeightAndRank
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Len(); got != tt.want {
				t.Errorf("ByHeightAndRank.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByHeightAndRank_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		b    ByHeightAndRank
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("ByHeightAndRank.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByHeightAndRank_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		b    ByHeightAndRank
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.Swap(tt.args.i, tt.args.j)
		})
	}
}
