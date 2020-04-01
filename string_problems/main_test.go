package main

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	tables := []struct {
		x string
		n bool
	}{
		{"abc", true},
	}

	for _, table := range tables {

		expected := table.n
		actual := IsUnique(table.x)
		if actual != expected {
			t.Errorf("Test of (%s) was incorrect, got: %v, want: %v.", table.x, actual, expected)
		}
	}
}
