package main

import "fmt"

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]

提示：
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案

进阶：你可以想出一个时间复杂度小于 O(n^2) 的算法吗？
*/
/*
思路：
一次遍历，如果target减去当前值没有，那么就放到哈希表里面，如果减去存在，就直接返回
*/

func twoSum1(nums []int, target int) []int {
	if nums == nil {
		return nil
	}

	hash := make(map[int]int)
	for index, val := range nums {
		if _, exists := hash[target-val]; exists {
			return []int{index, hash[target-val]}
		} else {
			hash[val] = index
		}
	}
	return nil
}

func main11111111111111() {
	a := []int{2, 7, 11, 15}
	//a := []int{3, 2, 4}
	//a := []int{3, 3}
	fmt.Println(twoSum1(a, 9))
}
