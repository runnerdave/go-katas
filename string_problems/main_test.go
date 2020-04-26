package main

import (
	"testing"
)

var (
	tables = []struct {
		x string
		n bool
	}{
		{"abc", true},
		{"abb", false},
		{"", true},
		{"♛", true},
		{"♛♛♛", false},
		{"", true},
		{" ", true},
		{"  ", false},
		{"111", false},
		{"123", true},
	}
)

// test if a string has all unique characters
func TestIsUnique(t *testing.T) {
	for _, table := range tables {
		expected := table.n
		actual := IsUnique(table.x)
		if actual != expected {
			t.Errorf("Test of (%s) was incorrect, got: %v, want: %v.", table.x, actual, expected)
		}
	}
}

// test if a string has all unique characters with loop implementation
func TestIsUniqueWithLoop(t *testing.T) {
	for _, table := range tables {
		expected := table.n
		actual := IsUniqueWithLoop(table.x)
		if actual != expected {
			t.Errorf("Test of (%s) was incorrect, got: %v, want: %v.", table.x, actual, expected)
		}
	}
}
