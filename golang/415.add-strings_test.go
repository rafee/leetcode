/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */

package golang

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
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
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkCarry(t *testing.T) {
	type args struct {
		digit byte
	}
	tests := []struct {
		name  string
		args  args
		want  byte
		want1 byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkCarry(tt.args.digit)
			if got != tt.want {
				t.Errorf("checkCarry() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkCarry() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
