package main

import (
	"reflect"
	"testing"
)

func TestGap(t *testing.T) {
	tables := []struct {
		x [3]int
		n []int
	}{
		{[3]int{0, 1, 0}, []int{}},
		{[3]int{0, 5, 2}, []int{3, 5}},
	}

	for _, table := range tables {

		expected := table.n
		actual := Gap(table.x[0], table.x[1], table.x[2])
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Gap of (%d) was incorrect, got: %d, want: %d.", table.x, actual, expected)
		}
	}
}

func TestGetPrimes(t *testing.T) {
	tables := []struct {
		x int
		n []int
	}{
		{2, []int{1, 2}},
		{5, []int{1, 2, 3, 5}},
		{20, []int{1, 2, 3, 5, 7, 11, 13, 17, 19}},
		{23, []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}},
	}

	for _, table := range tables {

		expected := table.n
		actual := GetPrimes(table.x)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("List of Primes for (%d) was incorrect, got: %d, want: %d.", table.x, actual, expected)
		}
	}
}
