package main

import "fmt"

/*
给你一个整数数组 nums，你需要确保数组中的元素 互不相同 。为此，你可以执行以下操作任意次：

从数组的开头移除 3 个元素。如果数组中元素少于 3 个，则移除所有剩余元素。
注意：空数组也视作为数组元素互不相同。返回使数组元素互不相同所需的 最少操作次数 。

示例 1：

输入： nums = [1,2,3,4,2,3,3,5,7]

输出： 2

解释：

第一次操作：移除前 3 个元素，数组变为 [4, 2, 3, 3, 5, 7]。
第二次操作：再次移除前 3 个元素，数组变为 [3, 5, 7]，此时数组中的元素互不相同。
因此，答案是 2。

示例 2：

输入： nums = [4,5,6,4,4]

输出： 2

解释：

第一次操作：移除前 3 个元素，数组变为 [4, 4]。
第二次操作：移除所有剩余元素，数组变为空。
因此，答案是 2。

示例 3：

输入： nums = [6,7,8,9]

输出： 0

解释：

数组中的元素已经互不相同，因此不需要进行任何操作，答案是 0。



提示：

1 <= nums.length <= 100
1 <= nums[i] <= 100
*/

/*
给你一个整数数组 nums，你需要确保数组中的元素 互不相同 。为此，你可以执行以下操作任意次：

从数组的开头移除 3 个元素。如果数组中元素少于 3 个，则移除所有剩余元素。
注意：空数组也视作为数组元素互不相同。返回使数组元素互不相同所需的 最少操作次数 。

题目要求：
确保数组中的元素不同，那么就是求  最长无重复后缀 ，因为是从数组开头移除的，所以，要求前面可以重复的，那么后续就是无重复， 所以，下标从后面开始
找到有重复的，那么 前面的就是 需要 全部移除的，每次操作三个元素，那么就是 个数/3， 不足3个元素按照三个元素计算。
*/

func minimumOperations(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}

	hashMap := make(map[int]struct{}, 0)
	for j := len(nums) - 1; j > 0; j-- {
		if _, exists := hashMap[nums[j]]; exists {
			return (j / 3) + 1
		}
		hashMap[nums[j]] = struct{}{}
	}

	return 0
}

func main() {
	a := []int{1, 2, 3, 4, 2, 3, 3, 5, 7}
	a = []int{5, 11, 14, 11}
	a = []int{3, 7, 7, 3}
	fmt.Println(minimumOperations(a))
}
