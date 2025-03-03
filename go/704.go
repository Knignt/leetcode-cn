package main

import (
	"fmt"
)

/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
提示：

你可以假设 nums 中的所有元素是不重复的。
n 将在 [1, 10000]之间。
nums 的每个元素都将在 [-9999, 9999]之间。
*/

func search(nums []int, target int) int {
	index := -1
	if len(nums) == 0 || target > nums[len(nums)-1] {
		return -1
	}

	//对于一个数字的数组的处理
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}

	mid := len(nums)/2 - 1
	left := 0
	right := len(nums) - 1
	for left <= right { //左下标要一直小于右下标，这里我写错了，写成了mid 小于right
		if nums[mid] == target {
			index = mid
			break
		}

		//这里需要去除mid,因为mid已经比较过了,所以就不需要再次处理了
		if nums[mid] > target {
			right = mid - 1 //右下标
		} else {
			left = mid + 1
		}

		mid = (left + right) / 2 //取中间值,类似于求二者平均值
	}

	return index
}

func main704() {
	//a := []int{-1, 0, 3, 5, 9, 12}
	a := []int{5}
	index := search(a, 2)
	fmt.Println(index)
}
