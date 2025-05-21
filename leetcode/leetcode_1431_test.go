package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Leetcode1431(t *testing.T) {
	type args struct {
		candies      []int
		extraCandies int
	}
	tests := []struct {
		args args
		want []bool
	}{
		{
			args{candies: []int{4, 2, 1, 1, 2}, extraCandies: 1},
			[]bool{true, false, false, false, false},
		},
		{
			args{candies: []int{2, 3, 5, 1, 3}, extraCandies: 3},
			[]bool{true, true, true, false, true},
		},
		{
			args{candies: []int{12, 1, 12}, extraCandies: 10},
			[]bool{true, false, true},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := Leetcode1431(tt.args.candies, tt.args.extraCandies)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
