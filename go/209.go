package main

import (
	"fmt"
	"math"
)

/*
209.长度最小的子数组
力扣题目链接(opens new window)

给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
提示：


1 <= target <= 10^9
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^5
*/

/*
总体思路：
滑动窗口
从0开始计算,查看当前的tmpSum是否大约等于target，如果是的话，那么在for循环里面，计算最小的窗口长度，然后将窗口的开始下标往前挪,一直到值小于target
然后继续再在之前的tmpSum上面继续求和,再循环

这里为什么是O(n)，因为下面的for只是挪动窗口的开始下标,而不是重新开始遍历
*/

func minSubArrayLen(target int, nums []int) int {
	if target < 0 || len(nums) == 0 {
		return 0
	}

	i, j, tmpSum := 0, 0, 0
	minLen := math.MaxInt
	for ; j < len(nums); j++ {
		tmpSum += nums[j]
		for tmpSum >= target { //题目要求是 ≥ target 的最小值
			minLen = min(minLen, j-i+1)
			tmpSum -= nums[i]
			i++
		}
	}

	if minLen == math.MaxInt {
		minLen = 0
	}

	return minLen
}

func main209() {
	//a := []int{1, 1, 1, 1, 1, 1, 1, 1}		11 返回0
	//a := []int{1, 2, 3, 4, 5}
	a := []int{10, 2, 3}
	fmt.Print(minSubArrayLen(6, a))
}
