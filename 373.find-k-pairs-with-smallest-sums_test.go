/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{[]int{1, 7, 11}, []int{2, 4, 6}, 3},
			want: [][]int{{1, 2}, {1, 4}, {1, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSmallestPairs(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		m    minHeap
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Len(); got != tt.want {
				t.Errorf("minHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minHeap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		m    minHeap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("minHeap.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minHeap_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		m    minHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_minHeap_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		m    *minHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Push(tt.args.x)
		})
	}
}

func Test_minHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		m    *minHeap
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
