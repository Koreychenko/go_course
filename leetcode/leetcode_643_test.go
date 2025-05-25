package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode643(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		args args
		want float64
	}{
		{
			args{nums: []int{1, 12, -5, -6, 50, 3}, k: 4},
			12.7500,
		},
		{
			args{nums: []int{5}, k: 1},
			5.0000,
		},
		{
			args{nums: []int{0, 1, 1, 3, 3}, k: 4},
			2.0000,
		},
		{
			args{nums: []int{4, 0, 4, 3, 3}, k: 5},
			2.8000,
		},
		{
			args{nums: []int{9, 7, 3, 5, 6, 2, 0, 8, 1, 9}, k: 6},
			5.333333333333333,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.FormatFloat(tt.want, 'f', -1, 64), func(t *testing.T) {
			if got := Leetcode643(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
