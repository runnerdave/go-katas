package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hallo du schöne Welt!")
}

// Prime Gap function, g and m are the limits and n is the target
func Gap(g, m, n int) []int {
	return []int{1, 1}
}

// GetPrimes, get prime numbers up to the number given
func GetPrimes(n int) []int {
	var primes []int
	m := make(map[int]bool)
	for i := 1; i <= n; i++ {
		m[i] = true
	}
	for i := 2; i <= n; i++ {
		for j := i; j <= n; j++ {
			if j%i == 0 && j != i {
				m[j] = false
			}
		}
		if i*i > n {
			break
		}
	}
	for k, v := range m {
		if v {
			primes = append(primes, k)
		}
	}
	sort.Ints(primes)
	return primes
}
