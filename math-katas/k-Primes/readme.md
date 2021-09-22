# K primes

from: https://www.codewars.com/kata/5726f813c8dcebf5ed000a6b/train/go

A natural number is called k-prime if it has exactly k prime factors, counted with multiplicity. For example:

````
k = 2  -->  4, 6, 9, 10, 14, 15, 21, 22, ...
k = 3  -->  8, 12, 18, 20, 27, 28, 30, ...
k = 5  -->  32, 48, 72, 80, 108, 112, ...
A natural number is thus prime if and only if it is 1-prime.
````

## Task:
Complete the function count_Kprimes (or countKprimes, count-K-primes, kPrimes) which is given parameters k, start, end (or nd) and returns an array (or a list or a string depending on the language - see "Solution" and "Sample Tests") of the k-primes between start (inclusive) and end (inclusive).

### Example:
countKprimes(5, 500, 600) --> [500, 520, 552, 567, 588, 592, 594]
Notes:

The first function would have been better named: findKprimes or kPrimes :-)
In C some helper functions are given (see declarations in 'Solution').
For Go: nil slice is expected when there are no k-primes between start and end.

## Lessons learned

Read the instructions properly and note that k-primes may or may not include multiplicity, 
that is 2 x 2 x 3 = 12 is seen here as a k 3 prime number but other definitions may consider it a k 2

## Installation

This kata uses ginkgo: https://github.com/onsi/ginkgo#getting-started

steps to install:
 - run these commands:
````
    go get -u github.com/onsi/ginkgo/ginkgo
	go get -u github.com/onsi/gomega/...
````
 - bootstrap the suite
````
	ginkgo bootstrap
````
 - run ginkgo
````
	ginkgo
````
