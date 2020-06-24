package main

import (
	"fmt"
	"runtime"
)

// Switch
// A switch statement is a shorter way to write a sequence if - else statements.
// It runs the first case whose value is equal to the condition expression.
// Go's switch is like the one in C, C++, Java, Javascript and PHP.
// Except that Go only runs the selected case, not all the cases that follow.
// In effect, the break statement that is needed at the end of each case in those languages is provided
// automatically in GO.
// Another important difference is that Go's switch cases need not to be constants,
// and the values involved need not be integers.

func main() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		//   freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
