package main

import (
	"sort"
	"strings"
)

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
