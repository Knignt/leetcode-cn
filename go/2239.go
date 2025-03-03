package main

import "fmt"

/*
给你一个长度为 n 的整数数组 nums ，请你返回 nums 中最 接近 0 的数字。如果有多个答案，请你返回它们中的 最大值 。

示例 1：

输入：nums = [-4,-2,1,4,8]
输出：1
解释：
-4 到 0 的距离为 |-4| = 4 。
-2 到 0 的距离为 |-2| = 2 。
1 到 0 的距离为 |1| = 1 。
4 到 0 的距离为 |4| = 4 。
8 到 0 的距离为 |8| = 8 。
所以，数组中距离 0 最近的数字为 1 。
示例 2：

输入：nums = [2,-1,1]
输出：1
解释：1 和 -1 都是距离 0 最近的数字，所以返回较大值 1 。

提示：

1 <= n <= 1000
-105 <= nums[i] <= 105
*/
func findClosestNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxNum := nums[0]
	var minLen int
	if nums[0] > 0 {
		minLen = nums[0]
	} else {
		minLen = nums[0] * -1
	}

	for _, num := range nums {
		tmp := num
		if num < 0 {
			tmp = num * -1
		}

		//如果距离相等,选择最大值，距离最小,直接无脑赋值
		if (num > maxNum && tmp == minLen) || tmp < minLen {
			maxNum = num
			minLen = tmp
		}
	}

	return maxNum
}

func main2239() {
	//a := []int{-4, -2, 1, 4, 8}
	a := []int{2, 1, 1, -1, 100000}
	b := findClosestNumber(a)
	fmt.Println(b)
}
