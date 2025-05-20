package main

import (
	"fmt"
)

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。


示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

提示：
n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/

/*
方法一：
1、从左往右、从右往左分别计算两个方向的“投影”的最大值
2、求出从左往右和从右往左的这两个值的最小值，然后和 height 进行相减，求和即可。
为什么取最小值，因为表示当前左边或者右边的最小值，这样不至于雨滴漏下去
*/

func trap(height []int) int {
	if len(height) <= 2 { //小于3的都没办法进行雨水的累积,会从任意一边溜出去
		return 0
	}

	ans := 0
	left := 0
	right := len(height) - 1
	preMax, sufMax := 0, 0

	//从左右两边开始,分别求出两边的最大值
	//在对于两边的最大值，求最小值
	//这里为什么是双指针
	//对于一个桶而言，决定可以存多少水是由其最短的木板来决定
	//如果他的左木板更短,则以这个为基准来计算,如果其右木板更短，则按照右木板来计算
	for left <= right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])
		if preMax < sufMax { //这里为什么要加的是left,如果left短，肯定加左边，不然水会流出去
			ans += preMax - height[left]
			left++
		} else {
			ans += sufMax - height[right]
			right--
		}
	}

	return ans
}

/*func trap(height []int) int {
	if len(height) <= 0 {
		return 0
	}

	//prefix
	prefix := make([]int, len(height))
	prefix[0] = height[0]
	maxNum := height[0]
	for i := 1; i < len(height); i++ {
		prefix[i] = max(maxNum, height[i])
		maxNum = prefix[i]
	}

	//subfix
	subfix := make([]int, len(height))
	subfix[len(height)-1] = height[len(height)-1]
	maxNum = subfix[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		subfix[i] = max(maxNum, height[i])
		maxNum = subfix[i]
	}

	ans := 0
	for i := 0; i < len(height); i++ {
		tmp := min(prefix[i], subfix[i])
		tmp = tmp - height[i]
		ans += tmp
	}

	return ans
}*/

func main() {
	a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	a = []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(a))
}
