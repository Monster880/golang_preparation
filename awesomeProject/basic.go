package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variable() {
	var a int
	var s int
	fmt.Printf("%d %q\n", a, s)
}

func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func variableInitialValue() {
	var a int = 3
	var s string = "abc"
	fmt.Println(a, s)
}

func triangle() {
	var a, b int = 3, 4
	c := math.Sqrt(float64(a*a + b*b))
	fmt.Println(a, b, c)
}

func consts() {
	const filename = "abc.txt"
	const a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func enums() {
	const (
		cpp        = 0
		java       = 1
		python     = 2
		golang     = 3
		javascript = 4
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

//func main() {
//	fmt.Println("Hello world")
//	variable()
//	variableInitialValue()
//	euler()
//	triangle()
//	consts()
//	enums()
//}
