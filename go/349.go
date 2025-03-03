package main

import "fmt"

/*
给定两个数组 nums1 和 nums2 ，返回 它们的
交集
 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的


提示：

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
*/

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	hash := make(map[int]int)
	var result = make(map[int]struct{})

	for i := 0; i < len(nums1); i++ {
		hash[nums1[i]]++
	}

	for j := 0; j < len(nums2); j++ {
		if hash[nums2[j]] > 0 {
			delete(result, nums2[j])
			result[nums2[j]] = struct{}{}
		}
	}

	var res []int
	for k, _ := range result {
		res = append(res, k)
	}

	return res
}

func main349() {
	a := []int{1, 2, 2, 1}
	b := []int{2, 2, 2, 2, 2}
	fmt.Println(intersection(a, b))
}
