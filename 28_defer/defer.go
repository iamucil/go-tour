package main

import (
	"fmt"
)

// Defer
// A defer statement defers the execution of a function until the surrounding functions returns.
// The deferred call's arguments are evaluated immediately,
// but the function call is not executed until the surrounding function returns.
func main() {
	defer fmt.Println("world")

	fmt.Println("01_hello")
}
