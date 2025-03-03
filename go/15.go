package main

import (
	"fmt"
	"sort"
)

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。
请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。


提示：

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func threeSum(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	var result [][]int

	sort.Ints(nums)

	right := len(nums) - 1
	for i := 0; i < right; i++ {
		if nums[i] > 0 && i == 0 { //如果第一项大于0，直接退出
			break
		}
		if i > 0 && nums[i] == nums[i-1] { //对于值相同的直接略过
			continue
		}

		left := i + 1
		for left < right {
			tmpSum := nums[i] + nums[left] + nums[right]
			if tmpSum == 0 {
				//append
				result = append(result, []int{nums[i], nums[left], nums[right]})

				//往后或者前挪动指针
				for left < right && nums[right-1] == nums[right] { //right 要和他的下一个比较
					right--
				}

				for left < right && nums[left+1] == nums[left] { //left要和他的上一个比较
					left++
				}
				//上面找到的是下一个或者上一个和当前的值不一样的，所以需要进行++和--，不然还是原来的值，容易导致死循环
				left++
				right--
			} else if tmpSum > 0 {
				right--
			} else {
				left++
			}
		}
		right = len(nums) - 1
	}

	return result
}

func main() {
	a := []int{-1, 0, 1, 2, -1, -4} //-4 -1 -1 0 1 2
	//a := []int{0, 0, 0, 0, 0, 0}
	//a := []int{1, -1, -1, 0}
	fmt.Println(threeSum(a))
}
