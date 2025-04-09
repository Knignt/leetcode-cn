package main

/*
给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，
请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
你所设计的解决方案必须只使用常量级的额外空间。


示例 1：
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

示例 2：
输入：numbers = [2,3,4], target = 6
输出：[1,3]
解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3] 。

示例 3：
输入：numbers = [-1,0], target = -1
输出：[1,2]
解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。


提示：
2 <= numbers.length <= 3 * 104
-1000 <= numbers[i] <= 1000
numbers 按 非递减顺序 排列
-1000 <= target <= 1000
仅存在一个有效答案
*/

/*
相向双指针
这里有一个很重要的前提： 数组是非递减顺序，也就是 递增 序列！！！如果不是递增序列，是不ok的
i 从0开始, j 要从最后一个开始，
如果两个值大于target表示当前的j所在的值过大，需要j--；
如果两个值小于target表示当前的i所在的值过小，需要i++

*/

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return make([]int, 0)
	}

	i := 0
	j := len(numbers) - 1

	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}

		if numbers[i]+numbers[j] < target {
			i++
			continue
		}

		if numbers[i]+numbers[j] > target {
			j--
		}
	}
	return make([]int, 0)
}
