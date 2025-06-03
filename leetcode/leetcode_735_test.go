package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Leetcode735(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args{nums: []int{5, 10, -5}},
			[]int{5, 10},
		},
		{
			args{nums: []int{8, -8}},
			[]int{},
		},
		{
			args{nums: []int{10, 2, -5}},
			[]int{10},
		},
		{
			args{nums: []int{-2, -1, 1, 2}},
			[]int{-2, -1, 1, 2},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode735(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
