/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

package golang

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{2, [][]int{{0, 1}, {1, 0}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextCourse(t *testing.T) {
	type args struct {
		graph   [][]int
		visited map[int]bool
		node    int
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
			if got := nextCourse(tt.args.graph, tt.args.visited, tt.args.node); got != tt.want {
				t.Errorf("nextCourse() = %v, want %v", got, tt.want)
			}
		})
	}
}
