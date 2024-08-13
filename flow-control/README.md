# Flow Control

## Functions

### Short Statements in Functions
- Go allows short statements before execution in functions, which is syntactic sugar.
- Variables declared in these short statements are only in scope until the end of the if statement.
- Short statement variables can be used elsewhere throughout the if statement.

Example:
```go
func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g \n", v, lim)
    }
    return lim
}
```

### Printing Formats
- Use `%g` for floats and `%d` for integers in `fmt.Printf()`.

## For Loops

- Go only has one looping construct: the for loop.
- Structure of a for loop:
  - Init statement: executed before the first iteration
  - Condition expression: evaluated before every iteration
  - Post statement: executed at the end of every iteration
- The loop continues until the condition statement is false.
- No parentheses are needed, just braces around the body.

Example:
```go
for i := 0; i < 10; i++ {
    sum += 1
}
```

- Init and post statements are optional.
- A conditional statement is always needed.
- For loops can also act as while loops.
- `for {}` will run forever.

## If Statements

- If statements don't need parentheses, just braces around the body.

## Switch Statements

- Switch in Go is similar to TypeScript, C, and C++.
- It only runs the selected case, not all cases that follow.
- The break statement needed in other languages is built into Go.
- Go switch cases don't need to be constants.
- Switch cases are evaluated from top to bottom.
- A switch without a condition is equivalent to `switch true`.

Example:
```go
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    fmt.Printf("%s.\n", os)
}
```

## Defer Statements

- A defer statement defers the execution of a function until the surrounding function returns.
- Deferred functions are pushed onto a stack.
- They are executed in Last In First Out (LIFO) order.

Example:
```go
defer fmt.Println("world")
fmt.Println("hello")
```

## Miscellaneous Notes

- Time in the Go Playground always appears to start at 11-10-2009, which is when Go was first publicly announced by Google (Go's birthday).