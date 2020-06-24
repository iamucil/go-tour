package main

// Function continued
// When two or more consecutive function parametes share a type, you can omit type from all but the last.
// In this example, we shortened
//	x int, y int
// to
//	x, y int

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
