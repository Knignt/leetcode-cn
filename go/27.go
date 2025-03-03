package main

import (
	"fmt"
)

/*
27. 移除元素
力扣题目链接(opens new window)

给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。 你不需要考虑数组中超出新长度后面的元素。

示例 2: 给定 nums = [0,1,2,2,3,0,4,2], val = 2, 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。

你不需要考虑数组中超出新长度后面的元素。
*/

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	swapFunc := func(i *int, j *int) {
		tmp := *i
		*i = *j
		*j = tmp
	}

	fast := 0
	slow := 0
	arrsLen := 0
	for ; fast < len(nums); fast++ { //快慢指针,快指针走在前面,快慢指针之间的是 和val相等的值
		if nums[fast] != val {
			swapFunc(&nums[fast], &nums[slow])
			slow++
			arrsLen++
		}
	}

	fmt.Println(nums)

	return arrsLen
}

func main() {
	a := []int{2, 3, 2, 3}
	cnt := removeElement(a, 3)
	fmt.Println(cnt)
}
