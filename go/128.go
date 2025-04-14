package main

import "fmt"

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

示例 2：
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

示例 3：
输入：nums = [1,0,1,2]
输出：3

提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
*/
func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	hashMap := make(map[int]bool, 0) //当前值:初始值
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]] = true
	}

	result := 0
	for num := range hashMap {
		if !hashMap[num-1] { //如果当前点为起点
			currentNum := num
			currentLen := 1

			for hashMap[currentNum+1] { //找到最长的序列
				currentLen++
				currentNum++
			}

			if currentLen > result { //比结果值大则更新
				result = currentLen
			}
		}
	}

	return result
}

func main() {
	a := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(a))
}
