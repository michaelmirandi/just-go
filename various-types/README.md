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


Understanding pointers and structs in Go is crucial for efficient memory management and data organization. While Go provides explicit control over value and reference semantics, other languages like Python and JavaScript have different approaches to object handling and memory management.