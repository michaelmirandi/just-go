# Packages, Variables, & Functions

## Packages

- The package name is the last element of the import path.
- Example: `math/rand` is package `rand`.

## Functions

- Function declarations go outside of the main file.
- Functions can take zero or more arguments.
- You can define multiple return types.

### Elegant Type Definitions
```go
func add_elegant(x, y int) int {
    return x + y
}
```
Consecutive parameter types can be defined together.

### Chained Functions
```go
func add_exotic(x, y int, z, k string) {
    fmt.Println(x+y, z, k)
}
```
Same syntax applies for chained functions with different types.

### Multiple Return Types
```go
func swap(x, y string) (string, string) {
    return y, x
}
```

### Naked Return
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```
Naked returns should only be used in short functions.

## Variables

- `var` declares a list of variables.
- Variables can be declared at package or function level.
- Outside a function, every statement must begin with a keyword (e.g., `var`, `func`, `const`).
- Variables can be factored into blocks.

### Short Assignment
```go
k := 3
```
Short assignment statements can be used in place of var.

### Multiple Variable Declaration
```go
var c, python, java = true, false, "no!"
```
Types are only needed if you don't initialize. If you initialize, the compiler will infer the type.

## Basic Types

- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64, uintptr
- byte (alias for uint8)
- rune (alias for int32, represents a Unicode code point)
- float32, float64
- complex64, complex128

### Zero Values
- 0 for numeric types
- false for boolean type
- "" (empty string) for strings

## Type Conversions

In Go, assignment between items of different types requires an explicit conversion:

```go
i := 42
f := float64(i)
u := uint(f)
```

## Type Inference

When the right hand side of the declaration is typed, the new variable is of that same type:

```go
var i int
j := i // j is an int

i := 42        // int
f := 3.142     // float64
g := 0.867 + .5i // complex128
```

Inference happens based on the precision of the constant.

## Constants

- Declared like variables but with the `const` keyword.
- Can be char, string, boolean, or numeric.
- CANNOT be declared using `:=` syntax.
- Numeric constants are high-precision values.
- An untyped constant takes the type needed by its context.

```go
const Big = 1 << 100
const Small = Big >> 99
```

## Miscellaneous Notes

- Exports begin with a capital letter.
- `:=` can be used to unpack multiple values from a function.
- Autoformat in Go is very effective ("SICK" as per the comment).