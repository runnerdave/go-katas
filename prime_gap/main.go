package main

import (
	"math"
	"sort"
)

// GetPrimes, get prime numbers up to the number given (sieve of Erathostenes)
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

// GetPrimesInRange, get prime numbers up to the number given
func GetPrimesInRange(o, n int) []int {
	var primes []int
	primes = GetPrimes(n)
	var init int
	for i, v := range primes {
		if v >= o {
			init = i
			break
		}
	}
	return primes[init:]
}

// Gap calculates the nearest target gap between primes
func Gap(g, m, n int) []int {
	primes := GetPrimesInRange(m, n)
	sort.Sort(sort.Reverse(sort.IntSlice(primes)))
	var result []int
	for i := 0; i < len(primes)-1; i++ {
		if primes[i]-primes[i+1] == g {
			result = []int{0, 0}
			result[0] = primes[i+1]
			result[1] = primes[i]
		}
	}
	return result
}

func GapDivisionCheck(g, m, n int) []int {
	var result []int
	for i := m; i <= n; i++ {
		if IsPrimeSqrt(i) {
			if i+g <= n && IsPrimeSqrt(i+g) {
				s := true
				for j := i + 1; j < g+i; j++ {
					if IsPrimeSqrt(j) {
						s = false
						break
					}
				}
				if s {
					result = []int{i, i + g}
					break
				}
			}

		}
	}
	return result
}

func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
