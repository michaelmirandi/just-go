package main

import (
	"fmt"
)

// a struct is a collection of fields
type Vertex struct {
	X, Y int
}

var (
	// can list subset of fields by using the attr: syntax (order is irrelevant)
	v1 = Vertex{1, 2} // type Vertex
	v2 = Vertex{X: 1} // Y: 0 is implicit
	v3 = Vertex{}     // X: 0, Y: 0 is implicit
	z  = &Vertex{1, 2}
)

func main() {
	// pointers hold the memory address of a value.
	i, j := 42, 2701
	//var emptyPointer *int
	// & generates a pointer to its operand
	// assigns the address of x to p
	p := &i
	// * denotes the pointer's underlying value
	fmt.Println(*p)
	// *p is a pointer to a value. zero value is nil
	// also known as dereferencing or indirecting
	// unlike c, go has no pointer arithmetic
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
	// * to access a value
	// & to assign an address in memory

	// pointers are helpful because
	// they allow passing large data structures by reference, instead of copying
	// mutability: functions can modify variables from the caller's scope
	// implementing data structures
	// allow for optional values
	// methods on value types: enable you to modify the receiver in methods

	// Go:
	// This gives you flexibility to explicitly pass by value or by reference (using pointers)
	// Basic types and structs are passed by value, slices/maps/channels by reference

	// Python:
	// All variables are references to objects
	// Reassignment of immutable types (int, string, tuple) creates new objects
	// Mutable objects (list, dict) can be modified in-place

	// TypeScript/JavaScript:
	// Primitives are passed by value
	// Objects (including arrays and functions) are passed by reference

	// General:
	// Value passing limits modifications to local function scope
	// Reference passing allows modifications to affect the original object
	// This can lead to unexpected behavior if not careful with mutable objects in Python and TS/JS

	v := Vertex{1, 2}
	// struct fields are accessed using a dot
	v.X = 4
	fmt.Println(v.X)

	// these fields can be accessed through struct pointer
	// go permits pointers to the object, then explicit references to the field
	z := &v
	// no need for explicit dereference (*)
	z.X = 1e9

	fmt.Println(v)

	// z is a pointer to a struct value
	fmt.Println(v1, z, v2, v3)

	/// b1

	// [n]T is an array of n values of type T
	// arrays are of fixed sizes
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// a slice is a dynamically sized- flexible view into the elements of an array
	// needs a low & high bound, separated by a colon
	var s []int = primes[1:4]
	fmt.Println(s)

	// a slice doesn't store any data.. just describes
	// if you change the value of a slice, it changes the underlying array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	first_part := names[0:2]
	second_part := names[1:3]

	fmt.Println(first_part, second_part)

	second_part[0] = "XXX"

	fmt.Println(first_part, second_part)
	fmt.Println(names)

	// a slice literal is alike an array literal without the length
	// creates an array, then a slice that references it...
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	array_of_structs := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(array_of_structs)

	// slice defaults
	// you can omit the high / low bounds to use the defaults instead
	// default is 0 for low & the lenght of the slice for the high

	temp_s := []int{2, 3, 5, 7, 11, 13}
	temp_s = s[1:4]
	fmt.Println(temp_s)

	temp_s = s[:2]
	fmt.Println(temp_s)

	temp_s = s[1:]
	fmt.Println(temp_s)

}
