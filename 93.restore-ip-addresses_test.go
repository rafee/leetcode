/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test 1",
			args: args{"25525511135"},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "Test 2",
			args: args{"1111"},
			want: []string{"1.1.1.1"},
		},
		{
			name: "Test 3",
			args: args{"010010"},
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_iPAdrressHelper(t *testing.T) {
	type args struct {
		s        string
		segments []string
		result   *[]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iPAdrressHelper(tt.args.s, tt.args.segments, tt.args.result)
		})
	}
}

func Test_chkNum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chkNum(tt.args.s); got != tt.want {
				t.Errorf("chkNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
