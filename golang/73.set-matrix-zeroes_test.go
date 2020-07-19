/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
		})
	}
}

func Test_setZero(t *testing.T) {
	type args struct {
		X      int
		Y      int
		col    []bool
		matrix [][]int
		row    []bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZero(tt.args.X, tt.args.Y, tt.args.col, tt.args.matrix, tt.args.row)
		})
	}
}

func Test_checkZero(t *testing.T) {
	type args struct {
		X      int
		Y      int
		matrix [][]int
	}
	tests := []struct {
		name  string
		args  args
		want  []bool
		want1 []bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkZero(tt.args.X, tt.args.Y, tt.args.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkZero() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("checkZero() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
