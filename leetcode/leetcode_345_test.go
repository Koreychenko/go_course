package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode345(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{
			"IceCreAm",
			"AceCreIm",
		},
		{
			"aI",
			"Ia",
		},
		{
			"aaI",
			"Iaa",
		},
		{
			"aacI",
			"Iaca",
		},
		{
			"a",
			"a",
		},
		{
			"c",
			"c",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode345(tt.args); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
