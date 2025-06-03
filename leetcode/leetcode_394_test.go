package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode394(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{
			"3[a]2[bc]",
			"aaabcbc",
		},
		{
			"3[a2[c]]",
			"accaccacc",
		},
		{
			"2[abc]3[cd]ef",
			"abcabccdcdcdef",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode394(tt.args); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
