package main

import (
	"testing"
)

var (
	uniqueData = []struct {
		x string
		n bool
	}{
		{"abc", true},
		{"abb", false},
		{"♛", true},
		{"♛♛♛", false},
		{"", true},
		{" ", true},
		{"  ", false},
		{"111", false},
		{"123", true},
	}

	permutationData = []struct {
		x [2]string
		n bool
	}{
		{[2]string{"abc", "cba"}, true},
		{[2]string{"aaa", "cba"}, false},
		{[2]string{"aaa", "aaa"}, true},
		{[2]string{"♛♛♛", "♛♛♛"}, true},
		{[2]string{"♛22", "22♛"}, true},
		{[2]string{"amor", "roma"}, true},
		{[2]string{"abba", "baba"}, true},
		{[2]string{"aba", "baba"}, false},
	}
)

// test if a string has all unique characters
func TestIsUnique(t *testing.T) {
	for _, table := range uniqueData {
		expected := table.n
		actual := IsUnique(table.x)
		if actual != expected {
			t.Errorf("Test of (%s) was incorrect, got: %v, want: %v.", table.x, actual, expected)
		}
	}
}

// test if a string has all unique characters with loop implementation
func TestIsUniqueWithLoop(t *testing.T) {
	for _, table := range uniqueData {
		expected := table.n
		actual := IsUniqueWithLoop(table.x)
		if actual != expected {
			t.Errorf("Test of (%s) was incorrect, got: %v, want: %v.", table.x, actual, expected)
		}
	}
}

// test if a string is a permutation of the other
func TestIsPermutation(t *testing.T) {
	for _, table := range permutationData {
		expected := table.n
		actual := IsPermutation(table.x)
		if actual != expected {
			t.Errorf("Test of (%s) was incorrect, got: %v, want: %v.", table.x, actual, expected)
		}
	}
}
