/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{3, 9},
			want: [][]int{{6, 2, 1}, {5, 3, 1}, {4, 3, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combinationSum3Ref(t *testing.T) {
	type args struct {
		start int
		k     int
		n     int
		res   *[][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			combinationSum3Ref(tt.args.start, tt.args.k, tt.args.n, tt.args.res)
		})
	}
}
