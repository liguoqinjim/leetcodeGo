package lab0088

func Merge(nums1 []int, m int, nums2 []int, n int) {
	for num := m + n; num > 0; num-- {
		if m == 0 {
			nums1[num-1] = nums2[n-1]
			n -= 1
			continue
		} else if n == 0 {
			return
		}

		max := nums1[m-1]
		max2 := nums2[n-1]

		if max2 > max {
			max = max2
			n -= 1
		} else {
			m -= 1
		}

		nums1[num-1] = max
	}
}
