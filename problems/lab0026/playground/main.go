package main

import "log"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	test(nums)

	log.Println(nums)
}

func test(nums []int) {
	nums[2] = 0
}
