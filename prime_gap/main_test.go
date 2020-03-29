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
		{[3]int{2, 100, 110}, []int{101, 103}},
		{[3]int{4, 100, 110}, []int{103, 107}},
		{[3]int{6, 100, 110}, nil},
		{[3]int{8, 300, 400}, []int{359, 367}},
		{[3]int{2, 3, 50}, []int{3, 5}},
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

func TestGetPrimesInRange(t *testing.T) {
	tables := []struct {
		x []int
		n []int
	}{
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 5}, []int{2, 3, 5}},
		{[]int{7, 21}, []int{7, 11, 13, 17, 19}},
		{[]int{17, 26}, []int{17, 19, 23}},
	}

	for _, table := range tables {
		expected := table.n
		actual := GetPrimesInRange(table.x[0], table.x[1])
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("List of Primes for (%d) was incorrect, got: %d, want: %d.", table.x, actual, expected)
		}
	}
}
