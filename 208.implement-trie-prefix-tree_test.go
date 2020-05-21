/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

package leetcode

import (
	"reflect"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		this *Trie
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Insert(tt.args.word)
		})
	}
}

func TestTrie_Search(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		this *Trie
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Search(tt.args.word); got != tt.want {
				t.Errorf("Trie.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_StartsWith(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		this *Trie
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.StartsWith(tt.args.prefix); got != tt.want {
				t.Errorf("Trie.StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrieConstructor(t *testing.T) {
	tests := []struct {
		name string
		want Trie
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrieConstructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrieConstructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
