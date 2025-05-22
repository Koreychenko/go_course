package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode605(t *testing.T) {
	type args struct {
		arg1 []int
		arg2 int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args{arg1: []int{1, 0, 0, 0, 1}, arg2: 1},
			true,
		},
		{
			args{arg1: []int{1, 0, 0, 0, 1}, arg2: 2},
			false,
		},
		{
			args{arg1: []int{1, 0, 0, 0, 0, 0, 1}, arg2: 2},
			true,
		},
		{
			args{arg1: []int{1, 0, 1, 0, 1, 0, 1}, arg2: 1},
			false,
		},
		{
			args{arg1: []int{1, 0, 0, 0, 0}, arg2: 2},
			true,
		},
		{
			args{arg1: []int{0, 0, 0, 0}, arg2: 3},
			false,
		},
		{
			args{arg1: []int{0, 0, 0, 0}, arg2: 2},
			true,
		},
		{
			args{arg1: []int{0, 0, 0, 0, 0}, arg2: 4},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode605(tt.args.arg1, tt.args.arg2); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
