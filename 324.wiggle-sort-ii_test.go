/*
 * @lc app=leetcode id=324 lang=golang
 *
 * [324] Wiggle Sort II
 */

package leetcode

import "testing"

func Test_wiggleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.args.nums)
		})
	}
}
