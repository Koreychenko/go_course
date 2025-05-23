package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode443(t *testing.T) {
	tests := []struct {
		args []byte
		want int
	}{
		{
			[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			6,
		},
		{
			[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			4,
		},
		{
			[]byte{'a'},
			1,
		},
		{
			[]byte{'a', 'a'},
			2,
		},
		{
			[]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
			6,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode443(tt.args); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
