package main

import (
	"reflect"
	"testing"
)

var (
	minDiffData = []struct {
		a []int32
		n int32
	}{
		{[]int32{1, 3, 3, 2, 4}, 3},
		{[]int32{5, 1, 3, 7, 3}, 6},
		{[]int32{3, 2}, 1},
		{[]int32{2, 3}, 1},
		{[]int32{2, 2147483647, 5}, 2147483645}, //max value for int32
	}

	intConvertData = []struct {
		a []int32
		n []int
	}{
		{[]int32{1, 3, 3, 2, 4}, []int{1, 3, 3, 2, 4}},
	}

	int32ConvertData = []struct {
		a []int
		n []int32
	}{
		{[]int{1, 3, 3, 2, 4}, []int32{1, 3, 3, 2, 4}},
	}
)

func TestMinDiff(t *testing.T) {
	for _, table := range minDiffData {
		expected := table.n
		actual := minDiff(table.a)
		if actual != expected {
			t.Errorf("Test of (%v) was incorrect, got: %v, want: %v.", table.a, actual, expected)
		}
	}
}

func TestConvertIntArray(t *testing.T) {
	for _, table := range intConvertData {
		expected := table.n
		actual := convertToIntArr(table.a)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Test of (%v) was incorrect, got: %v, want: %v.", table.a, actual, expected)
		}
	}
}

func TestConvertInt32Array(t *testing.T) {
	for _, table := range int32ConvertData {
		expected := table.n
		actual := convertToInt32Arr(table.a)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Test of (%v) was incorrect, got: %v, want: %v.", table.a, actual, expected)
		}
	}
}
