/*
 * @lc app=leetcode id=332 lang=golang
 *
 * [332] Reconstruct Itinerary
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_findItinerary(t *testing.T) {
	type args struct {
		tickets [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// {
		// 	name: "Test 1",
		// 	args: args{[][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}},
		// 	want: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		// },
		// {
		// 	name: "Test 2",
		// 	args: args{[][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}},
		// 	want: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		// },
		// {
		// 	name: "Test 3",
		// 	args: args{[][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}},
		// 	want: []string{"JFK", "NRT", "JFK", "KUL"},
		// },
		{
			name: "Test 4",
			args: args{[][]string{{"EZE", "TIA"}, {"EZE", "HBA"}, {"AXA", "TIA"}, {"JFK", "AXA"}, {"ANU", "JFK"}, {"ADL", "ANU"}, {"TIA", "AUA"}, {"ANU", "AUA"}, {"ADL", "EZE"}, {"ADL", "EZE"}, {"EZE", "ADL"}, {"AXA", "EZE"}, {"AUA", "AXA"}, {"JFK", "AXA"}, {"AXA", "AUA"}, {"AUA", "ADL"}, {"ANU", "EZE"}, {"TIA", "ADL"}, {"EZE", "ANU"}, {"AUA", "ANU"}}},
			want: []string{"JFK", "AXA", "AUA", "ADL", "ANU", "AUA", "ANU", "EZE", "ADL", "EZE", "ANU", "JFK", "AXA", "EZE", "TIA", "AUA", "AXA", "TIA", "ADL", "EZE", "HBA"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findItinerary(tt.args.tickets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItinerary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_visitAirport(t *testing.T) {
	type args struct {
		src     string
		destMap map[string][]string
		pre     []string
		ttl     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := visitAirport(tt.args.src, tt.args.destMap, tt.args.pre, tt.args.ttl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("visitAirport() = %v, want %v", got, tt.want)
			}
		})
	}
}
