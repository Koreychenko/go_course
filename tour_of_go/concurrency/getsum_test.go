package main

import "testing"

type testcase struct {
	numbers []int
	want    int
}

func TestGetSum(t *testing.T) {
	for _, tc := range getTestCases() {
		t.Run("tc", func(t *testing.T) {
			got := getSum(tc.numbers)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

func getTestCases() []testcase {
	return []testcase{
		{[]int{7, 2, 8, -9, 4, 0}, 12},
		{[]int{7, 2, 8, -9, 4, 1}, 13},
		{[]int{7, 2, 8, -9, 4}, 12},
	}
}
