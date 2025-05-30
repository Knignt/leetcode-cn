package main

import (
	"fmt"
	"math"
)

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组中的一个连续部分。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

示例 2：
输入：nums = [1]
输出：1

示例 3：
输入：nums = [5,4,-1,7,8]
输出：23

提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104

进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
*/
func maxSubArray(nums []int) int {

	if len(nums) <= 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	subSum := make([]int, len(nums)+1)
	subSum[0] = 0
	for i := 0; i < len(nums); i++ {
		subSum[i+1] = nums[i] + subSum[i]
	}

	currentMin := subSum[0] //前缀和最小值
	currentSumMax := math.MinInt
	for i := 1; i < len(subSum); i++ {
		current := subSum[i] - currentMin
		if current > currentSumMax {
			currentSumMax = current
		}

		if subSum[i] < currentMin {
			currentMin = subSum[i]
		}
	}

	//fmt.Println(currentSumMax, currentMin)

	return currentSumMax

}

func main() {

	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//a = []int{5, 4, -1, 7, 8}
	a = []int{-2, 1}
	fmt.Println(maxSubArray(a))
}
