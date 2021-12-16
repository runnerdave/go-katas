package main

import (
	"fmt"
	"strings"
)

func staircase(n int32) {
	for i := int(n); i > 0; i-- {
		for j := 1; j < int(n)+1; j++ {
			if j < i {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		if i > 0 {
			fmt.Print("\n")
		}
	}
}

func staircaseOptimized(n int32) {
	for i := int32(1); i <= n; i++ {
		fmt.Printf("%*s\n", n, strings.Repeat("#", int(i)))
	}
}
