package main

import (
	"fmt"
	"sort"
)

/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

示例 1：
输入：nums = [1,2,0]
输出：3
解释：范围 [1,2] 中的数字都在数组中。

示例 2：
输入：nums = [3,4,-1,1]
输出：2
解释：1 在数组中，但 2 没有。

示例 3：
输入：nums = [7,8,9,11,12]
输出：1
解释：最小的正数 1 没有出现。

提示：
1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
*/
func firstMissingPositive(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}

	//对于只有单个数字的数组的处理
	if len(nums) == 1 {
		if nums[0] == 1 {
			return 2
		} else {
			return 1
		}
	}

	//这里不是O(n^2)的时间复杂度，是O(n),这里是为了一直替换到这个位置，比如1放在0位置上,2放在1位置上
	n := len(nums)
	for i := 0; i < n; i++ {
		//这里对于小于0的不做处理
		//要判断下标是否存在在当前数组里面，如果不存在，则不处理
		//最关键的是已交换的数据不再做交换
		//确保nums[i]要放到 nums[nums[i]-1]的位置去,并且之前已交换的不做交换.
		for nums[i] > 0 && nums[i]-1 <= n && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func firstMissingPositive1(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == 1 {
			return 2
		} else {
			return 1
		}
	}

	sort.Ints(nums)
	if nums[0] > 0 && nums[0] != 1 {
		return 1
	}

	for i := 0; i < len(nums); i++ {

		if nums[i] > 0 {

			if i > 0 && nums[i-1] <= 0 && nums[i] >= 0 && nums[i] != 1 {
				return 1
			}

			if i < len(nums)-1 && nums[i+1]-nums[i] > 1 {
				return nums[i] + 1
			}
		}
	}

	if nums[len(nums)-1] < 0 {
		return 1
	} else {
		return nums[len(nums)-1] + 1
	}
}

func main() {

	a := []int{0, -1, 3, 1} //2
	//a = []int{3, 4, -1, 1}  //2
	a = []int{1, 2, 0} //3
	//a = []int{-1, -2, -60, 40, 43} //1
	//a = []int{1}                   //2
	//a = []int{1, 1000}             //2

	//a = []int{0, 3} //1
	a = []int{2, 2} //1
	a = []int{1, 1}
	a = []int{7, 8, 9, 11, 12}
	fmt.Println(firstMissingPositive(a))

}
