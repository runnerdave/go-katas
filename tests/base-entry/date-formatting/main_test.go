package main

import (
	"testing"
)

var (
	dates = []struct {
		x      string
		target string
	}{
		{"2nd June 1976", "1976-06-02"},
		{"1st January 2003", "2003-01-01"},
		{"3rd October 1050", "1050-10-03"},
	}
)

func TestFormatDate(t *testing.T) {
	for _, table := range dates {
		expected := table.target
		actual := formatDate(table.x)
		if actual != expected {
			t.Errorf("Test of (%s) was incorrect, got: %v, want: %v.", table.x, actual, expected)
		}
	}
}
