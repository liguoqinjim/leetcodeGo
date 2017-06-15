package main

import (
	"fmt"
	"math"
)

func main() {
	reverse(3231)

	fmt.Println("max=", math.MaxInt32)
	fmt.Println("min=", math.MinInt32)
}

func reverse(x int) int {
	fmt.Println("x=", x)
	if x == 0 {
		return 0
	}

	i := 10
	num := make([]int, 0)
	for {
		rem := x % i
		num = append(num, rem/(i/10))
		if rem == x {
			break
		}
		i *= 10
	}
	result := 0
	fmt.Println(num)
	for n, v := range num {
		result += v * int(math.Pow(10, float64(len(num)-1-n)))
	}

	if result == math.MaxInt32 {
		return 0
	}

	return result
}
