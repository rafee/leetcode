/*
 * @lc app=leetcode id=1835 lang=golang
 *
 * [1835] Find XOR Sum of All Pairs Bitwise AND
 */

package golang

import "testing"

func Test_getXORSum(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getXORSum(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("getXORSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
