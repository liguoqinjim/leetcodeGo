package lib

func twoSum(nums []int, target int) []int {
	as := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				as[0] = i
				as[1] = j
			}
		}
	}

	return as
}

