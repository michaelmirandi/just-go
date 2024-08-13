package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// prints using default formats
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// short statement before execution, syntactic sugar
	// variables declared in the short statement are only in scope until the end of the if
	if v := math.Pow(x, n); v < lim {
		return v
	} else { // can also use short statement variables elsewhere throughout if statement
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}

// estimating sqrt of an input float64 x
// newton's method
func Sqrt(x float64) float64 {
	z := 1.0
	iteration_count := 0
	for i := 0; true; i++ {
		fmt.Println("HELLO")
		zNew := z - (z*z-x)/(2*z)
		fmt.Println(zNew)
		if math.Abs(zNew-z) < 1e-8 {
			break // how to break out of a loop
		}
		z = zNew
		// %g for float, %d for int
		fmt.Printf("z: %g, i: %d \n", z, i)
		iteration_count = i
	}
	fmt.Println("iterations ", iteration_count)
	fmt.Println("our sqrt: ", z)
	fmt.Println("actual sqrt: ", math.Sqrt(x))
	return z
}

func main() {
	// go only has 1 looping construct: for loop
	sum := 0
	// init statement: executed before first iteration
	// condition expression: evaluated before every iteration
	// post statement: executed at end of every iteration
	// for loops through until the condition statement is false
	// no parenthesis, just braces around body
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	// init & post statements are optional
	// always need a conditional statement
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	// for loops also act as while

	for sum < 10000 {
		sum += sum
	}
	fmt.Println(sum)

	// for {} will run forever...

	// if statments also don't need parenthesis, just braces around body
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 5),
		pow(3, 2, 20),
	)

	Sqrt(158234.681357522)

	// switch in go is like TS, C, C++
	// except it only runs the selected case, not all cases that follow
	// break statement that is needed in other languages above
	// is baked into Go
	// go switch cases don't need to be constants
	fmt.Print("Go runs on... ")
	// another short statement...
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// switch cases evaluation from top to bottom
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	// time in go playground always appears to start at 11-10-2009...
	// when go was first publicly announced by google
	// go birthday

	// switch without a conditoin is just true...
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Buongiorno")
	case t.Hour() > 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Buonasera")
	}

	// defer statement defers the execution of a function until the surrouding function returns
	defer fmt.Println("world")
	fmt.Println("hello")

	// deferred functions are pushed to a stack
	// lifo execution
	fmt.Println("couting...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
