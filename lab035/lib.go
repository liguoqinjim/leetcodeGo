package lab035

func searchInsert(nums []int, target int) int {
	for n, v := range nums {
		if target <= v {
			return n
		}
	}
	return len(nums)
}
