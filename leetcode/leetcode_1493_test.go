package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode1493(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{nums: []int{1, 1, 0, 1}},
			3,
		},
		{
			args{nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}},
			5,
		},
		{
			args{nums: []int{1, 1, 1}},
			2,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode1493(tt.args.nums); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
