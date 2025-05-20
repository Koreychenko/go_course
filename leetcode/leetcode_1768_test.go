package leetcode

import "testing"

func Test_Leetcode1768(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{word1: "abc", word2: "pqr"},
			"apbqcr",
		},
		{
			args{word1: "ab", word2: "pqrs"},
			"apbqrs",
		},
		{
			args{word1: "abcd", word2: "pq"},
			"apbqcd",
		},
		{
			args{word1: "a", word2: "b"},
			"ab",
		},
		{
			args{word1: "a", word2: "bc"},
			"abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := Leetcode1768(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
