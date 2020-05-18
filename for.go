package main

import "fmt"

// For
// Go has only one looping construct, the for loop.
// The basic for loop has three components separated by semicolons:
//
//	The init statement: executed before the first iteration
//	The condition expression: evaluated before every iteration
//	The post statement: executed at the end of every iteration
//
// The init statement will often be a short variable declaration, and the variables
// declared there are visible only in the scope of the for statement.
//
// The loop will stop iterating once the boolean condition evaluated to false.
// Note: unline other languages like C, Java, or JavaScript there are no patentheses surrounding
// the three components of the for statement and the braces { } are always required.

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
