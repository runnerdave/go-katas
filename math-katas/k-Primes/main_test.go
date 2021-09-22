package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest1(k, start, nd int, exp []int) {
	var ans = CountKprimes(k, start, nd)
	//Expect(ans).To(Equal(exp))
	CheckIfAns(ans, exp)
}

func dotest2(s int, exp int) {
	var ans = Puzzle(s)
	Expect(ans).To(Equal(exp))
}

func dotest3(p int, exp []int) {
	var ans = PrimeFactors(p)
	Expect(ans).To(Equal(exp))
}

func dotest4(p int, exp map[int]int) {
	var ans = PrimeFactorsMap(p)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest1(5, 1000, 1100, []int{1020, 1026, 1032, 1044, 1050, 1053, 1064, 1072, 1092, 1100})
		dotest1(5, 500, 600, []int{500, 520, 552, 567, 588, 592, 594})
		dotest1(12, 100000, 100100, nil)
	})
	It("should handle basic puzzle", func() {
		dotest2(138, 1) //[2 + 8 + 128]
		dotest2(143, 2) //[3 + 12 + 128] and [7 + 8 + 128]
	})
	It("should handle prime factors", func() {
		dotest3(8, []int{2, 2, 2})
		dotest3(18, []int{2, 3, 3})
		dotest3(10, []int{2, 5})
		dotest3(30, []int{2, 3, 5})
	})
	It("should handle prime factors with map", func() {
		dotest4(8, map[int]int{2: 3})
		dotest4(18, map[int]int{2: 1, 3: 2})
		dotest4(10, map[int]int{2: 1, 5: 1})
		dotest4(30, map[int]int{2: 1, 3: 1, 5: 1})
	})
})
