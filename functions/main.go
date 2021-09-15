package main

import (
	"fmt"
)

var anonymous = func() {
	fmt.Println("hello world first class function")
}

var anonymousWithParam = func(n string) {
	fmt.Println("Welcome", n)
}

type add func(a, b int) int

func higherOrderFunctionAsParam(a func(a, b int) int) {
	fmt.Printf("Higher order, func as param value = %v\n", a(60, 7))
}

func higherOrderReturnFunction() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	a := anonymous
	a()
	fmt.Printf("\n%T\n", a)

	n := anonymousWithParam
	n("Gophers")

	func() {
		fmt.Println("true anonymous")
	}()

	var p add = func(a int, b int) int {
		return a + b
	}
	fmt.Printf("user defined function type: 2 + 2 = %v\n", p(2, 2))

	higherOrderFunctionAsParam(p)

	s := higherOrderReturnFunction()
	fmt.Printf("Higher order, return function value = %v\n", s(60, 7))

	c := 5
	func() {
		fmt.Println("closure: c =", c)
	}()

	d := appendStr()
	e := appendStr()
	fmt.Println(d("World"))
	fmt.Println(e("Everyone"))

	fmt.Println(d("Gopher"))
	fmt.Println(e("!"))

}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
