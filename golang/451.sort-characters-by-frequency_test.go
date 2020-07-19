/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

package golang

import (
	"reflect"
	"testing"
)

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{"cccaaa"},
			want: "cccaaa",
		},
		{
			name: "Test 2",
			args: args{"tree"},
			want: "eert",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); got != tt.want {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_freqPQ_Len(t *testing.T) {
	tests := []struct {
		name string
		pq   freqPQ
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Len(); got != tt.want {
				t.Errorf("freqPQ.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_freqPQ_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		pq   freqPQ
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("freqPQ.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_freqPQ_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		pq   freqPQ
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_freqPQ_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		pq   *freqPQ
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Push(tt.args.x)
		})
	}
}

func Test_freqPQ_Pop(t *testing.T) {
	tests := []struct {
		name string
		pq   *freqPQ
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("freqPQ.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_freqPQ_update(t *testing.T) {
	type args struct {
		item     *freqCount
		value    byte
		priority int
	}
	tests := []struct {
		name string
		pq   *freqPQ
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.update(tt.args.item, tt.args.value, tt.args.priority)
		})
	}
}
