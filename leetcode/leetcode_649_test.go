package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode649(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{
			"RD",
			"Radiant",
		},
		{
			"RDD",
			"Dire",
		},
		{
			"RDDD",
			"Dire",
		},
		{
			"RDDDR",
			"Dire",
		},
		{
			"RRDDR",
			"Radiant",
		},
		{
			"RRR",
			"Radiant",
		},
		{
			"DDRRR",
			"Dire",
		},
		{
			"DRDRR",
			"Dire",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode649(tt.args); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
