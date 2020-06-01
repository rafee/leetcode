/*
 * @lc app=leetcode id=821 lang=golang
 *
 * [821] Shortest Distance to a Character
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_shortestToChar(t *testing.T) {
	type args struct {
		S string
		C byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{"loveleetcode", 'e'},
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestToChar(tt.args.S, tt.args.C); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
