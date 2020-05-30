/*
 * @lc app=leetcode id=973 lang=golang
 *
 * [973] K Closest Points to Origin
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		K      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kClosestSelect(t *testing.T) {
	type args struct {
		points [][]int
		K      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosestSelect(tt.args.points, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosestSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSelect(t *testing.T) {
	type args struct {
		points [][]int
		start  int
		end    int
		K      int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSelect(tt.args.points, tt.args.start, tt.args.end, tt.args.K)
		})
	}
}

func Test_swapPoints(t *testing.T) {
	type args struct {
		i      int
		pivot  int
		points [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swapPoints(tt.args.i, tt.args.pivot, tt.args.points)
		})
	}
}

func Test_kClosestHeap(t *testing.T) {
	type args struct {
		points [][]int
		K      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosestHeap(tt.args.points, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosestHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_closePoints_Len(t *testing.T) {
	tests := []struct {
		name string
		c    closePoints
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("closePoints.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_closePoints_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		c    closePoints
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("closePoints.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_closePoints_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		c    closePoints
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_closePoints_Push(t *testing.T) {
	type args struct {
		p interface{}
	}
	tests := []struct {
		name string
		c    *closePoints
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Push(tt.args.p)
		})
	}
}

func Test_closePoints_Pop(t *testing.T) {
	tests := []struct {
		name string
		c    *closePoints
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closePoints.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
