package main

import (
	"reflect"
	"sort"
	"strings"
)

// IsUnique tests if string contains only unique strings
func IsUnique(s string) bool {
	a := strings.Split(s, "")
	result := true
	sort.Strings(a)
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			result = false
		}
	}
	return result
}

// IsUnique tests if string contains only unique strings using loops in the implementation
func IsUniqueWithLoop(s string) bool {
	a := strings.Split(s, "")
	result := true
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if i != j {
				if a[i] == a[j] {
					result = false
				}
			}
		}
	}
	return result
}

// IsPermutation tests if given two strings one is a permutation of the other
func IsPermutation(s [2]string) bool {
	m1 := mapStrings(s[0])
	m2 := mapStrings(s[1])
	return reflect.DeepEqual(m1, m2)
}

func mapStrings(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Split(s, "")
	for _, v := range a {
		m[v]++
	}
	return m
}
