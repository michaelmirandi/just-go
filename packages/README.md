# Packages, Variables, & Functions

## Table of Contents
1. [Package Declaration](#package-declaration)
2. [Imports](#imports)
3. [Functions](#functions)
4. [Variables](#variables)
5. [Basic Types](#basic-types)
6. [Type Conversions](#type-conversions)
7. [Type Inference](#type-inference)
8. [Constants](#constants)

## Package Declaration
- The `package` keyword declares which package the file belongs to.
- `package main` is used for executable programs.

## Imports
- Use the `import` keyword to include external packages.
- Multiple imports can be factored into parentheses.
- The last element of the import path is the package name (e.g., `math/rand` is package `rand`).

## Functions
- Function declaration: `func name(parameters) return_type { ... }`
- Multiple parameters of the same type: `func add(x, y int) int { ... }`
- Multiple return values: `func swap(x, y string) (string, string) { ... }`
- Naked return: `return` without arguments returns named return values.
- Exported functions start with a capital letter.

## Variables
- Declare with `var` keyword: `var name type = value`
- Short declaration inside functions: `name := value`
- Multiple declarations: `var x, y int = 1, 2` or `x, y := 1, 2`
- Factored variable declaration:
  ```go
  var (
      name1 type1 = value1
      name2 type2 = value2
  )
  ```
- Zero values: 0 for numeric types, false for booleans, "" for strings.

## Basic Types
- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64, uintptr
- byte (alias for uint8)
- rune (alias for int32, represents Unicode code point)
- float32, float64
- complex64, complex128

## Type Conversions
- Explicit conversion required: `T(v)` converts value v to type T
- Example: `float64(intValue)` or `uint(floatValue)`

## Type Inference
- The `:=` syntax allows type inference based on the right-hand side value.
- For constants, inference is based on precision of the constant.

## Constants
- Declared using `const` keyword: `const Name = value`
- Can be character, string, boolean, or numeric values.
- Cannot use `:=` syntax for constants.
- Numeric constants are high-precision values.
- Untyped constants take the type needed by their context.

## Additional Notes
- Use `fmt.Println()` for standard output.
- Use `fmt.Printf()` for formatted output.
- The `math/rand` package provides random number functionality.
- The `math` package provides mathematical constants and functions.