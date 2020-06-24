package main

import (
	"fmt"
	"math"
)

// If and else
// Variables declared inside an if short statement are also available inside of the else blocks.
// (Both calls to pow return their results before the call to fmt.Println in main begins)

func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %gv\n", v, limit)
	}
	// can't use v here, though
	return limit
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
