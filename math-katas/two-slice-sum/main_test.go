package main

import (
	"reflect"
	"testing"
)

var (
	data = []struct {
		nums   []int
		target int
		n      []int
		ok     bool
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}, true},
	}
)

func TestTwoSum(t *testing.T) {
	for _, table := range data {
		expected := table.n
		actual := TwoSumBruteForce(table.nums, table.target)
		if reflect.DeepEqual(actual, expected) != table.ok {
			t.Errorf("Test of (%v,%d) was incorrect, got: %v, want: %v.", table.nums, table.target, actual, expected)
		}
	}
}
