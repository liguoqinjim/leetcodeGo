package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func searchInsert(nums []int, target int) int {
	for n, v := range nums {
		if target <= v {
			return n
		}
	}
	return len(nums)
}
