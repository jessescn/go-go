package main

import (
	"fmt"
	"strconv"
)

var varOutsideFunction string = "Outside Var"

// We can declare multiple variables with same var declaration
var (
	variableOne   string = "One"
	variableTwo   string = "Two"
	variableThree string = "Three"
)

func main() {
	// we can declare variables in three different ways
	// 1. We can declare variable without initializing
	var k int
	k = 30
	// 2. We can declare with initial value, both ways we can assure the variable type
	var i float32 = 10.5
	// 3. We can let the compile infer the type
	j := 25

	// Using strconv to convert int to string
	var l string
	l = strconv.Itoa(k)

	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Println(l)
}
