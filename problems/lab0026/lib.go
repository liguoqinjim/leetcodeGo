package lab0026

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	num := 0
	j := nums[0] - 1
	for i := 0; i < len(nums); i++ {
		if j != nums[i] {
			j = nums[i]
			nums[num] = j
			num++
		}
	}

	return num
}
