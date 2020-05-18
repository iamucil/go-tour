package main

import "fmt"

// Type inference
// When declaring a variable without specifying an explicit
// (either by using the :=  syntax or var = expression syntax),
// The variable's type is infered from the value on the right hand side.
// When the right hand side of the declaration is type, the new variable
// is of that same type:
//
//		var i int
//		j := i // j is an int
//
// But when the right hand side contains an untyped numeric constant, the new variable
// may be an int, float64, or complex128 depending on the precision on the constant.
//
//		i := 42						// int
//		f := 3.142				// float64
//		g := 0.867 + 0.5i	// complex128
//
// Try chaning the initial value of v in the example code and observe how its type is affected.

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}
