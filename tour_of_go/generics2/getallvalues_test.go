package main

import (
	"slices"
	"testing"
)

func TestGetAllValues(t *testing.T) {
	nodes := prepareList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	got := nodes[7].GetAllNextValues()
	want := []int{8, 9, 10}
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
