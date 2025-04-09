package main

import "fmt"

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。

示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例 2：
输入：height = [1,1]
输出：1

提示：
n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/

func maxArea(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}

	i, j := 0, len(height)-1
	ans := 0
	for j > i {
		tmp := (j - i) * (min(height[i], height[j]))
		ans = max(ans, tmp)
		if height[i] > height[j] { //每次移动高度较小的指针。因为移动较大的指针会导致宽度减少，而高度受限于较小的那个，容量必然减少。移动较小的指针可能找到更高的垂线，从而有机会增加容量。
			j--
		} else {
			i++
		}
	}

	return ans
}

func main() {
	a := []int{8, 7, 2, 1}
	fmt.Println(maxArea(a))
}
