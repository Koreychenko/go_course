package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode1004(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, k: 2},
			6,
		},
		{
			args{nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, k: 3},
			10,
		},
		{
			args{nums: []int{0, 0, 0, 1}, k: 3},
			4,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode1004(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
