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
		{[]int{3, 2, 4}, 6, []int{1, 2}, true},
		{[]int{3, 3}, 6, []int{0, 1}, true},
	}
)

func TestTwoSum(t *testing.T) {
	for _, table := range data {
		expected := table.n
		actual := TwoSum(table.nums, table.target)
		if reflect.DeepEqual(actual, expected) != table.ok {
			t.Errorf("Test of (%v,%d) was incorrect, got: %v, want: %v.", table.nums, table.target, actual, expected)
		}
	}
}

func TestTwoSumBruteForce(t *testing.T) {
	for _, table := range data {
		expected := table.n
		actual := TwoSumBruteForce(table.nums, table.target)
		if reflect.DeepEqual(actual, expected) != table.ok {
			t.Errorf("Test of (%v,%d) was incorrect, got: %v, want: %v.", table.nums, table.target, actual, expected)
		}
	}
}
