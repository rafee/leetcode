/*
 * @lc app=leetcode id=468 lang=golang
 *
 * [468] Validate IP Address
 */

package leetcode

import "testing"

func Test_validIPAddress(t *testing.T) {
	type args struct {
		IP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{"256.256.256.256"},
			want: "Neither",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validIPAddress(tt.args.IP); got != tt.want {
				t.Errorf("validIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkIPv4(t *testing.T) {
	type args struct {
		segments []string
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
			if got := checkIPv4(tt.args.segments); got != tt.want {
				t.Errorf("checkIPv4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkIPv6(t *testing.T) {
	type args struct {
		segments []string
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
			if got := checkIPv6(tt.args.segments); got != tt.want {
				t.Errorf("checkIPv6() = %v, want %v", got, tt.want)
			}
		})
	}
}
