package main

func moveZeroes(nums []int) {
	zero_num, index_value := 0, 0
	for zero_num+index_value < len(nums) {
		if nums[index_value] == 0 {
			nums = append(nums[:index_value], nums[index_value+1:]...)
			nums = append(nums, 0)
			zero_num++
		} else {
			index_value++
		}
	}
}
