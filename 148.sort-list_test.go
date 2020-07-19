/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortListHelper(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name  string
		args  args
		want  *ListNode
		want1 *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := sortListHelper(tt.args.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortListHelper() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("sortListHelper() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
