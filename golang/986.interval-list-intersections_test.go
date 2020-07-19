/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_intervalIntersection(t *testing.T) {
	type args struct {
		A [][]int
		B [][]int
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
			if got := intervalIntersection(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInterval(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getInterval(tt.args.A, tt.args.B)
			if got != tt.want {
				t.Errorf("getInterval() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getInterval() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
