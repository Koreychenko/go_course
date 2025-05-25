package leetcode

import (
	"testing"
)

func Test_Leetcode1456(t *testing.T) {
	type args struct {
		s string
		k int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args{s: "abciiidef", k: 3},
			3,
		},
		{
			args{s: "aeiou", k: 2},
			2,
		},
		{
			args{s: "leetcode", k: 3},
			2,
		},
		{
			args{s: "a", k: 1},
			1,
		},
		{
			args{s: "ibpbhixfiouhdljnjfflpapptrxgcomvnb", k: 33},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.s, func(t *testing.T) {
			if got := Leetcode1456(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
