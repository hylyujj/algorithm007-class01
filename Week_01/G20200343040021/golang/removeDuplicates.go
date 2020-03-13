package main

func removeDuplicates(nums []int) int {
    num_value := 0
    for i := 0; i< len(nums); i++ {
	if nums[num_value] != nums[i] {
	    nums[num_value + 1] = nums[i]
	    num_value++
        }
    }
    return num_value + 1
}

func main() {
    aaa := []int{0,0,1,1,1,2,2,3,3,4}
    removeDuplicates(aaa)
}
