/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectTopKNums(t *testing.T) {
	type args struct {
		numFreq [][2]int
		low     int
		high    int
		k       int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectTopKNums(tt.args.numFreq, tt.args.low, tt.args.high, tt.args.k)
		})
	}
}

func Test_swapNums(t *testing.T) {
	type args struct {
		high    int
		low     int
		numFreq [][2]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swapNums(tt.args.high, tt.args.low, tt.args.numFreq)
		})
	}
}
