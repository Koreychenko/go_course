package leetcode

import (
	"strconv"
	"testing"
)

func Test_Leetcode1071(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{str1: "ABCABC", str2: "ABC"},
			"ABC",
		},
		{
			args{str1: "ABABAB", str2: "ABAB"},
			"AB",
		},
		{
			args{str1: "ABABABAB", str2: "ABAB"},
			"ABAB",
		},
		{
			args{str1: "LEET", str2: "CODE"},
			"",
		},
		{
			args{str1: "TAUXXTAUXXTAUXXTAUXXTAUXX", str2: "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"},
			"TAUXX",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := Leetcode1071(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
