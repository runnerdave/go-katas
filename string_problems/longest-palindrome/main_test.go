package main

import (
	"testing"
)

var (
	palindromicData = []struct {
		x      string
		target string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"aaa", "aaa"},
		{"bb", "bb"},
		{"bbbb", "bbbb"},
		{"a", "a"},
		{"ac", "a"},
	}

	palindromeData = []struct {
		x string
		n bool
	}{
		{"babab", true},
		{"babad", false},
		{"cbbd", false},
		{"aaa", true},
		{"a", true},
		{"ac", false},
	}
)

func TestFindLongestPalindrome(t *testing.T) {
	for _, table := range palindromicData {
		expected := table.target
		actual := findPalindrome(table.x)
		if actual != expected {
			t.Errorf("Test of (%s) was incorrect, got: %v, want: %v.", table.x, actual, expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	for _, table := range palindromeData {
		expected := table.n
		actual := isPalindrome(table.x)
		if actual != expected {
			t.Errorf("Test of (%s) was incorrect, got: %v, want: %v.", table.x, actual, expected)
		}
	}
}
