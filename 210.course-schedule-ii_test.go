/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

package leetcode

// func Test_findOrder(t *testing.T) {
// 	type args struct {
// 		numCourses    int
// 		prerequisites [][]int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []int
// 	}{
// 		{
// 			name: "Test 1",
// 			args: args{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
// 			want: []int{0, 1, 2, 3},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("findOrder() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_goDeep(t *testing.T) {
// 	type args struct {
// 		courseGraph [][]int
// 		visited     []bool
// 		i           int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := goDeep(tt.args.courseGraph, tt.args.visited, tt.args.i); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("goDeep() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
