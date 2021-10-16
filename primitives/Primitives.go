package main

import (
	"fmt"
)

func main() {
	fmt.Println("Operate with differents types")
	var n uint16 = 42
	fmt.Printf("%v, %T\n", n, n)

	// cannot operate with differents variables types, its necessary do a conversion in one of then
	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b))

	fmt.Println("\nLogical Operations with integers")
	// using integers in logical operation compare their binary values
	c := 10 // 1010
	d := 3  // 0011

	fmt.Println(c & d)  // 0010
	fmt.Println(c | d)  // 1011
	fmt.Println(c ^ d)  // 1001
	fmt.Println(c &^ d) // 0100

	fmt.Println("\nExponential Operations with >> and << ")
	e := 8              // 2^3
	fmt.Println(e << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(e >> 3) // 2^3 / 2^3 = 2^0

	fmt.Println("\nFloat variables")
	var f float64 = 3.14
	f = 13.7e72
	f = 2.1e14
	fmt.Printf("%v, %T\n", f, f)

	fmt.Println("\nComplex values")
	var g complex64 = 2i + 1
	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", real(g), real(g))
	fmt.Printf("%v, %T\n", imag(g), imag(g))

	fmt.Println("\nWorking with strings!")
	h := "String example"
	h2 := " Another string example"
	fmt.Printf("%v, %T\n", h+h2, h+h2)

	h3 := []byte(h)
	fmt.Printf("%v, %T\n", h3, h3)

	fmt.Println("\nRunes? Runes!")
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
