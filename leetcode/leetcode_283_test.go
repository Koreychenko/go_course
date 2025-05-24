package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Leetcode283(t *testing.T) {
	tests := []struct {
		args []int
		want []int
	}{
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
		{
			[]int{0},
			[]int{0},
		},
		{
			[]int{1, 0},
			[]int{1, 0},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{0, 1},
			[]int{1, 0},
		},
		{
			[]int{0, 1, 0},
			[]int{1, 0, 0},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode283(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
