package main

func plusOne(digits []int) []int {
	if digits[len(digits) - 1] != 9 {
		digits[len(digits) - 1] ++
		return digits
	} else {
		nine_num := 0
		for i:=len(digits) - 1; i>=0; i-- {
			// 从后往前遍历
			if digits[i] == 9 {
				//遇到元素为9的nine_num增加一个，并且把该位置的元素换成0
				digits[i] = 0
				nine_num++
			} else {
				//遇到第一个元素不为9的，当前元素+1
				digits[i]++
				break
			}
		}
		//9的个数更切片长度一样，说明需要扩容一个
		if nine_num == len(digits) {
			digits = append([]int{1}, digits...)
		}
		return digits
	}
}
