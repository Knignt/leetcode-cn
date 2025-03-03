package main

import (
	"container/list"
	"fmt"
)

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

示例 2：
输入：nums = [1], k = 1
输出：[1]

提示：
1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length
*/

// 单调队列
/*
设计单调队列的时候，pop，和push操作要保持如下规则：

pop(value)：如果窗口移除的元素value等于单调队列的出口元素，那么队列弹出元素，否则不用任何操作
push(value)：如果push的元素value大于入口元素的数值，那么就将队列入口的元素弹出，直到push元素的数值小于等于队列入口元素的数值为止
保持如上规则，每次窗口移动的时候，只要问que.front()就可以返回当前窗口的最大值。


1、如果push进去的值大于前面的值，将前面的值全部pop出去
2、如果当前的值小于列表的尾部，push进去
3、如果超过窗口，pop出去窗口第一个值,然后再获取第一个值
*/

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	ans := make([]int, 0, len(nums)-k+1) // 预分配空间
	deque := list.New()
	for i, val := range nums {
		//如果最新入队的比队列最后的元素大，那么要将队列一直popright到比他大这里
		for deque.Len() > 0 && val >= nums[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())
		}

		//如果当前对首的下标超过窗口，直接将之前的一直pop到满足窗口内的即可
		for deque.Len() > 0 && i-deque.Front().Value.(int) >= k {
			deque.Remove(deque.Front())
		}

		deque.PushBack(i)

		//获取当前窗口的最大值
		if i >= k-1 {
			ans = append(ans, nums[deque.Front().Value.(int)])
		}

	}
	return ans
}

func printContainerList(l list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d\t", e.Value)
	}
	fmt.Println()
}

func main239() {
	//nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//k := 3
	nums := []int{7, 2, 4}
	k := 2
	fmt.Println(maxSlidingWindow(nums, k))
}
