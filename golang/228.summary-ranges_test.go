/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRngStr(t *testing.T) {
	type args struct {
		num int
		pre int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRngStr(tt.args.num, tt.args.pre); got != tt.want {
				t.Errorf("getRngStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
