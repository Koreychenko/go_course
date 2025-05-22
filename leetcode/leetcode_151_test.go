package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode151(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{
			"the sky is blue",
			"blue is sky the",
		},
		{
			"  hello world  ",
			"world hello",
		},
		{
			"a good   example",
			"example good a",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode151(tt.args); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
