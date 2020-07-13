/*
 * @lc app=leetcode id=948 lang=golang
 *
 * [948] Bag of Tokens
 */

package leetcode

import "testing"

func Test_bagOfTokensScore(t *testing.T) {
	type args struct {
		tokens []int
		P      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[]int{71, 55, 82}, 54},
			want: 0,
		},
		{
			name: "Test 2",
			args: args{[]int{100, 200, 300, 400}, 200},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{[]int{54}, 150},
			want: 1,
		},
		{
			name: "Test 4",
			args: args{[]int{54}, 50},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bagOfTokensScore(tt.args.tokens, tt.args.P); got != tt.want {
				t.Errorf("bagOfTokensScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
