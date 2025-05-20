package main

import "fmt"

/*
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
请 不要使用除法，且在 O(n) 时间复杂度内完成此题。

示例 1:
输入: nums = [1,2,3,4]
输出: [24,12,8,6]

示例 2:
输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]

提示：
2 <= nums.length <= 105
-30 <= nums[i] <= 30
输入 保证 数组 answer[i] 在  32 位 整数范围内

进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？
（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
*/

func productExceptSelf(nums []int) []int {

	if len(nums) <= 0 {
		return nil
	}

	//left
	n := len(nums)
	left := make([]int, n)
	left[0] = nums[0]

	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i]
	}

	//right
	right := make([]int, n)
	right[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i]
	}

	//fmt.Println(left, right)

	for i := 1; i < n-1; i++ {
		nums[i] = left[i-1] * right[i+1]
	}

	nums[0] = right[1]
	nums[n-1] = left[n-2]

	return nums
}

func productExceptSelf1(nums []int) []int {

	if len(nums) <= 0 {
		return nil
	}

	num := 1
	numZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			numZero++
			continue
		} else {
			num *= nums[i]
		}
	}

	//对于全为0的处理
	if numZero == len(nums) && num == 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && numZero >= 1 {
			nums[i] = 0
		} else {
			if nums[i] == 0 && numZero == 1 {
				nums[i] = num
			} else if nums[i] == 0 && numZero > 1 {
				nums[i] = 0
			} else {
				nums[i] = num / nums[i]
			}
		}
	}

	return nums
}

func main() {

	a := []int{0, 0}
	a = []int{1, 2, 3, 4}
	a = []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(a))

}
