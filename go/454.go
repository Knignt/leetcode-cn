package main

import "fmt"

/*
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0


示例 1：

输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
输出：2
解释：
两个元组如下：
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
示例 2：

输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
输出：1


  提示：

n == nums1.length
n == nums2.length
n == nums3.length
n == nums4.length
1 <= n <= 200
-228 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 228
*/

/*
思路：
先把nums1 和 nums2 的 和 两次循环求出来，并在map中记录这个和出现了几次
再把nums3 和 nums4 的 和 在两次循环里面求出来，并且如果 这个和 的绝对值 在前面map出现过
就把次数和 之前map里面和出现的次数相加，即为结果
为什么这里是加上这个值呢？主要是因为表示这个值可以被选择两次,否则就会被少算
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	if nums1 == nil || nums2 == nil || nums3 == nil || nums4 == nil {
		return 0
	}

	aAndB := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			aAndB[nums1[i]+nums2[j]]++
		}
	}

	cnt := 0
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			if aAndB[0-nums3[i]-nums4[j]] > 0 {
				cnt += aAndB[0-nums3[i]-nums4[j]] //为什么这里是加上这个值呢？主要是因为表示这个值可以被选择两次,否则就会被少算
			}
		}
	}

	return cnt
}

func main454() {
	a := []int{-1, -1}
	b := []int{-1, 1}
	c := []int{-1, 1}
	d := []int{1, -1}
	fmt.Println(fourSumCount(a, b, c, d))
}
