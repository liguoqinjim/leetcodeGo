package main

import "fmt"

func main() {
	//printInt("(){}[]")

	//testSlice()

	fmt.Println(isValid("(){}[]"))
}

func testSlice() {
	a := []int{1}
	a = a[:len(a)-1]
	fmt.Println(a)
}

func printInt(s string) {
	for i := 0; i < len(s); i++ {
		a := s[i]
		fmt.Println(int(a))
	}
}

func isValid(s string) bool {
	left := make([]int, 1)
	left[0] = int(s[0])
	for i := 1; i < len(s); i++ {
		a := int(s[i])
		if len(left) == 0 {
			left = append(left, a)
			continue
		}
		if a-left[len(left)-1] != 1 && a-left[len(left)-1] != 2 {
			left = append(left, a)
		} else {
			left = left[:len(left)-1]
		}
		fmt.Println(left)
	}
	if len(left) == 0 {
		return true
	} else {
		return false
	}
}
