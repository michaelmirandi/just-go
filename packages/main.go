package main

// package name is the last element of import path
// ex. math/rand is package rand

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// function declaration goes outside of main file
// elegant type definitions
func add(x int, y int) int {
	return x + y
}

// consecutive parameter types can be defined as such:
func add_elegant(x, y int) int {
	return x + y
}

// same goes for chained functions... nice lil syntactic sugar
func add_exotic(x, y int, z, k string) {
	fmt.Println(x+y, z, k)
}

// can define multiple return types
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

// naked returns should only be used in short functions

// var declares a list of variables

// *** function level

// var c, python, java bool

var i, j int = 1, 2

// outside a function, every statement must begin w/ a keyword
// a := 5 // not valid

// variables can be factored into blocks

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// autoformat is SICK

func main() {
	// package level
	fmt.Println("My favorite number is", rand.Intn(10))
	// exports begin with a capital letter
	fmt.Println("Magical: ", math.Pi)

	// functions take zero or more arguments
	fmt.Println("My first function add(x, y): ", add(rand.Intn(100), rand.Intn(1000)))
	fmt.Println("My first elegant, shared type function add_elegant(x, y): ", add_elegant(rand.Intn(100), rand.Intn(1000)))
	fmt.Println("My first exotic, shared type function add_exotic(x, y): ")
	add_exotic(rand.Intn(100), rand.Intn(1000), "big", "mike")
	// := to unpack multiple values from a function
	y, x := swap("me", "fuck")
	fmt.Println("Multi-return type", y, x)
	x_1, y_1 := split(17)
	fmt.Println("Naked return function results: ", x_1, y_1)

	// functions can be declared at the package or function level
	// var i int
	var c, python, java = true, false, "no!" // types are only needed if you don't initialize
	// if you intialize, the compiler will give it a value.
	fmt.Println(i, c, python, java)

	// short assignment statements can be used in place of var
	k := 3

	fmt.Println("short assignment: ", k)

	// Basic Types

	// bool
	// string
	// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
	// byte // alias for uint8
	// rune // alias for int32, represents a unicode code point
	// float32 float64
	// complex64 complex128

	fmt.Printf("Type: %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value %v\n", z, z)

	// zero value is 0 for numeric types
	// 	false for boolean type
	// 	"" for strings
	// if not initialized, then these values will appear

	var t int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %v", t, f, b, s)
}
