/*
 * @lc app=leetcode id=212 lang=golang
 *
 * [212] Word Search II
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test 1",
			args: args{[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
				[]string{"oath", "pea", "eat", "rain"},
			},
			want: []string{"oath", "eat"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWord(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want wdObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWord(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wdObject_Len(t *testing.T) {
	tests := []struct {
		name string
		o    wdObject
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Len(); got != tt.want {
				t.Errorf("wdObject.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
