package main

func rotate(nums []int, k int) {
	if k % len(nums) == 0 {
		return
	}

	step_value := k % len(nums)
	is_head, heads_index, mid_value, next_index := 0, 0, 0, 0

	for i := 0; i < len(nums); i++ {
		if is_head == 0 {
			next_index = step_value + heads_index
			nums[next_index], mid_value = nums[heads_index], nums[next_index]
			is_head++
		} else {
			nums[next_index], mid_value = mid_value, nums[next_index]
			is_head = 1
			if next_index == heads_index {
				is_head = 0
				heads_index++
			}
		}
		next_index = (next_index + step_value) % len(nums)
	}
}

func main() {
    aaa := [1,2,3,4,5,6]
    k := 2
    rotate(aaa, k)
}
