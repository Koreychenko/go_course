package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Leetcode238(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
		{
			[]int{0, 0},
			[]int{0, 0},
		},
		{
			[]int{0, 0, 1},
			[]int{0, 0, 0},
		},
		{
			[]int{0, 4, 0},
			[]int{0, 0, 0},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode238(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
