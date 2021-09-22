package common

import (
	"fmt"
	"math"
)

//produces the Sieve of Erathostenes
func Sieve(n int) []int {
	c := make([]int, n)
	for i := range c {
		c[i] = 1
	}
	for i := 2; i*i <= n; i++ {
		if c[i] == 1 {
			for p := i * 2; p < n; p += i {
				c[p] = 0
			}
		}
	}
	return c
}

//k-primes without multiplicity
// ie: 2, 14, 18 gives: []int{14, 15, 18} -- 7*2, 5*3, (3*3)*2
func CountKPrimes(k, start, nd int) []int {
	c := make([]int, nd+1)
	fmt.Println(c)
	for i := 2; i <= nd; i++ {
		if c[i] == 0 {
			for j := i; j <= nd; j += i {
				c[j]++
			}
		}
	}
	fmt.Println(c)
	fmt.Println(k)
	var r []int
	for i := start; i <= nd; i++ {
		if c[i] == k {
			r = append(r, i)
		}
	}
	fmt.Println(r)
	return r
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
