package main

import (
	"reflect"
	"testing"
)

func TestGap(t *testing.T) {
	tables := []struct {
		x int
		n []int
	}{
		{0, []int{}},
		{3, []int{2, 2}},
	}

	for _, table := range tables {

		expected := table.n
		actual := Gap(table.x)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Gap of (%d) was incorrect, got: %d, want: %d.", table.x, actual, expected)
		}
	}
}
