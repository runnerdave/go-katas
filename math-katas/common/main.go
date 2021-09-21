package common

import "fmt"

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
