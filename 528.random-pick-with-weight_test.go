/*
 * @lc app=leetcode id=528 lang=golang
 *
 * [528] Random Pick with Weight
 */

package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		w []int
	}
	tests := []struct {
		name string
		args args
		want Solution
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_PickIndex(t *testing.T) {
	tests := []struct {
		name string
		this *Solution
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.PickIndex(); got != tt.want {
				t.Errorf("Solution.PickIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_bSearch(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		this *Solution
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.bSearch(tt.args.target); got != tt.want {
				t.Errorf("Solution.bSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
