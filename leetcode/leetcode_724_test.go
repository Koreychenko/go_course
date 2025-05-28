package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Leetcode724(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{
			[]int{1, 7, 3, 6, 5, 6},
			3,
		},
		{
			[]int{1, 2, 3},
			-1,
		},
		{
			[]int{2, 1, -1},
			0,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode724(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
