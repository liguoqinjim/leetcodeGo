package main

import (
	"fmt"
	"math"
)

func main() {
	var a = 0
	a = math.MaxInt64
	fmt.Printf("%d\n", a)
	a += 1
	fmt.Printf("%d\n", a)
}
