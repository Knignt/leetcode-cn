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
	//a = []int{1, 2, 0}      //3
	//a = []int{-1, -2, -60, 40, 43} //1
	//a = []int{1}                   //2
	//a = []int{1, 1000}             //2

	a = []int{0, 3} //1
	a = []int{2, 2} //1
	fmt.Println(firstMissingPositive(a))

}
