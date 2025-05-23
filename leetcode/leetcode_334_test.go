package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode334(t *testing.T) {
	tests := []struct {
		args []int
		want bool
	}{
		{
			[]int{2, 1, 5, 0, 4, 6},
			true,
		},
		{
			[]int{5, 4, 3, 2, 1},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5},
			true,
		},
		{
			[]int{20, 100, 10, 12, 5, 13},
			true,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode334(tt.args); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
