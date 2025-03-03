package main

import "fmt"

/*
977.有序数组的平方
力扣题目链接(opens new window)

给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
#算法公开课

*/

func sortedSquares(nums []int) []int {
	if len(nums) <= 0 {
		return []int{0}
	}
	i := 0
	j := len(nums) - 1
	result := make([]int, len(nums))
	lenRes := len(result) - 1

	//双指针法,一个指针在最左边，一个在最右边，最左边和最右边的值平方比较,大的值放在result的后面,result的指针左移即可
	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			result[lenRes] = nums[i] * nums[i]
			i++
		} else {
			result[lenRes] = nums[j] * nums[j]
			j--
		}
		lenRes--
	}

	return result
}

func main977() {
	a := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(a))
}
