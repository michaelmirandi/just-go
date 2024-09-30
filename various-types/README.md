# Various Types

## Structs

In Go, a struct is a collection of fields. It's used to group related data together.

```go
type Vertex struct {
    X, Y int
}
```

This defines a `Vertex` struct with two integer fields, `X` and `Y`.

## Pointers

Pointers hold the memory address of a value.

```go
i, j := 42, 2701
p := &i  // p is a pointer to i
fmt.Println(*p)  // prints the value of i
*p = 21  // sets i through the pointer p
fmt.Println(i)  // prints 21
```

- `&` operator generates a pointer to its operand.
- `*` operator denotes the pointer's underlying value (dereferencing).

### Example: Modifying a Variable Through a Pointer

```go
p = &j
*p = *p / 37
fmt.Println(j)  // j is now changed
```

This example shows how to modify a variable indirectly through a pointer.

## Struct Initialization

There are multiple ways to initialize a struct:

```go
v1 := Vertex{1, 2}  // Specify all fields in order
v2 := Vertex{X: 1}  // Specify fields by name (Y:0 is implicit)
v3 := Vertex{}      // All fields are set to zero values
z := &Vertex{1, 2}  // Create a pointer to a new struct instance
```

## Accessing Struct Fields

Struct fields are accessed using a dot notation:

```go
v := Vertex{1, 2}
v.X = 4
fmt.Println(v.X)  // Prints 4
```

When using a pointer to a struct, you can access fields directly:

```go
z := &v
z.X = 1e9  // Equivalent to (*z).X = 1e9
```

Go automatically dereferences the pointer when accessing struct fields.

# Go Arrays and Slices

This README provides an overview and examples of working with arrays and slices in Go.

## Arrays

In Go, an array is a fixed-size sequence of elements of the same type. The size is part of the array's type, making arrays in Go more rigid but also more memory-efficient than in some other languages.

### Declaring and Initializing Arrays

You can declare and initialize arrays in several ways:

```go
// Declare an array of two strings
var a [2]string
a[0] = "Hello"
a[1] = "World"

// Declare and initialize in one line
primes := [6]int{2, 3, 5, 7, 11, 13}
```

### Examples

```go
fmt.Println(a[0], a[1])  // Output: Hello World
fmt.Println(a)           // Output: [Hello World]
fmt.Println(primes)      // Output: [2 3 5 7 11 13]
```

## Slices

A slice is a dynamically-sized, flexible view into the elements of an array. It's more common than arrays in Go.

### Creating Slices

You can create a slice by specifying a low and high bound, separated by a colon:

```go
var s []int = primes[1:4]
```

### Slice Properties

- Slices don't store data, they just describe a section of an underlying array.
- Changes to a slice modify the corresponding elements in its underlying array.
- Other slices that share the same underlying array will see those modifications.

### Examples

```go
names := [4]string{"John", "Paul", "George", "Ringo"}

first_part := names[0:2]
second_part := names[1:3]

fmt.Println(first_part, second_part)  // Output: [John Paul] [Paul George]

second_part[0] = "XXX"
fmt.Println(names)  // Output: [John XXX George Ringo]
```

### Slice Literals

A slice literal is like an array literal without the length. It creates an array, then builds a slice that references it:

```go
q := []int{2, 3, 5, 7, 11, 13}
r := []bool{true, false, true, true, false, true}
```

You can also create slices of structs:

```go
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
```

### Slice Defaults

When slicing, you may omit the high or low bounds to use their defaults. The default is zero for the low bound and the length of the slice for the high bound.

```go
s := []int{2, 3, 5, 7, 11, 13}

fmt.Println(s[1:4])  // Output: [3 5 7]
fmt.Println(s[:2])   // Output: [2 3]
fmt.Println(s[1:])   // Output: [3 5 7 11 13]
```

## Language Comparisons

### Go
- Explicitly pass by value or by reference (using pointers)
- Basic types and structs are passed by value
- Slices, maps, and channels are passed by reference

### Python
- All variables are references to objects
- Reassignment of immutable types creates new objects
- Mutable objects can be modified in-place

### TypeScript/JavaScript
- Primitives are passed by value
- Objects (including arrays and functions) are passed by reference

### General Observations
- Value passing limits modifications to local function scope
- Reference passing allows modifications to affect the original object
- This can lead to unexpected behavior if not careful with mutable objects in Python and TS/JS