package main

func twoSum(nums []int, target int) []int {
    num_map := make( map[int]int)
    result_list := make([]int, 2)
    for index_value, first_num := range nums {
        second_num := target - first_num
        if _, ok := num_map[second_num]; ok {
            result_list[0] = num_map[second_num]
            result_list[1] = index_value
            return result_list
        }
        num_map[first_num] = index_value
    }
    return result_list
}
