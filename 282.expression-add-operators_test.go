/*
 * @lc app=leetcode id=282 lang=golang
 *
 * [282] Expression Add Operators
 */

package golang

import (
	"reflect"
	"testing"
)

func TestOperation_String(t *testing.T) {
	tests := []struct {
		name string
		o    Operation
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("Operation.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addOperators(t *testing.T) {
	type args struct {
		num    string
		target int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test 1",
			args: args{"1", 1},
			want: []string{"1"},
		},
		{
			name: "Test 2",
			args: args{"12", 2},
			want: []string{"1*2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOperators(tt.args.num, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOperators() = %v, want %v", got, tt.want)
			}
		})
	}
}
