package main

import (
	"math"
)

func CountKprimes(k, start, nd int) []int {
	var m []map[int]int
	for i := start; i <= nd; i++ {
		m = append(m, PrimeFactorsMap(i))
	}
	var r []int
	for i, n := range m {
		var v int
		for _, value := range n {
			v += value
		}
		if v == k {
			r = append(r, start+i)
		}
	}
	return r
}

func Puzzle(s int) int {
	sol := 0
	k1 := CountKprimes(1, 1, s)
	k3 := CountKprimes(3, 1, s)
	k7 := CountKprimes(7, 1, s)
	for _, a := range k1 {
		for _, b := range k3 {
			for _, c := range k7 {
				if s == a+b+c {
					sol++
				}
			}
		}
	}
	return sol
}

func PrimeFactors(n int) []int {
	var p []int
	for n%2 == 0 {
		p = append(p, 2)
		n /= 2
	}
	for i := 3; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		for n%i == 0 {
			n /= i
			p = append(p, i)
		}
	}
	if n > 2 {
		p = append(p, n)
	}
	return p
}

func PrimeFactorsMap(n int) map[int]int {
	p := make(map[int]int)
	for n%2 == 0 {
		p[2]++
		n /= 2
	}
	for i := 3; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		for n%i == 0 {
			n /= i
			p[i]++
		}
	}
	if n > 2 {
		p[n]++
	}
	return p
}
