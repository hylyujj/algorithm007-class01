package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	end_1, end_2 := m-1, n-1
	sum := m+n-1
	for end_1 >= 0 && end_2 >= 0 {
		if nums2[end_2] >= nums1[end_1] {
			nums1[sum] = nums2[end_2]
			end_2--
			sum--
		} else {
			nums1[sum] = nums1[end_1]
			end_1--
			sum--
		}
	}
	for ; end_2 >= 0; end_2-- {
		nums1[sum] = nums2[end_2]
		sum--
	}
	return
}

