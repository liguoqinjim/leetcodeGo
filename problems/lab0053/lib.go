package lab0053

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	max := nums[0]
	for i := len(nums); i >= 1; i-- {
		for j := 0; j+i-1 < len(nums); j += 1 {
			temp := 0
			for k := j; k < j+i; k++ {
				temp += nums[k]
			}
			if temp > max {
				max = temp
			}
		}
	}

	return max
}

//动态规划
func MaxSubArray2(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
