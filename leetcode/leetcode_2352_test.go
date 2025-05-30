package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode2352(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{grid: [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}},
			1,
		},
		{
			args{grid: [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}},
			3,
		},
		{
			args{grid: [][]int{{13, 13}, {13, 13}}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.want), func(t *testing.T) {
			if got := Leetcode2352(tt.args.grid); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
