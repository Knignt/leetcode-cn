package main

/*
1343. 大小为 K 且平均值大于等于阈值的子数组数目
中等
相关标签
相关企业
提示
给你一个整数数组 arr 和两个整数 k 和 threshold 。
请你返回长度为 k 且平均值大于等于 threshold 的子数组数目。

示例 1：
输入：arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
输出：3
解释：子数组 [2,5,5],[5,5,5] 和 [5,5,8] 的平均值分别为 4，5 和 6 。其他长度为 3 的子数组的平均值都小于 4 （threshold 的值)。

示例 2：
输入：arr = [11,13,17,23,29,31,7,5,2,3], k = 3, threshold = 5
输出：6
解释：前 6 个长度为 3 的子数组平均值都大于 5 。注意平均值不是整数
*/

// 小于滑动窗口的往里面加，大于则判断值是否大于和，再将左边往前挪
func numOfSubarrays1(arr []int, k int, threshold int) int {
	if len(arr) == 0 {
		return 0
	}
	result := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i < k-1 {
			continue
		}

		if sum >= threshold*k {
			result++
		}
		sum -= arr[i-k+1]
	}

	return result
}

// 滑动窗口的做法
func numOfSubarrays(arr []int, k int, threshold int) int {
	if len(arr) == 0 {
		return 0
	}
	result := 0

	i, j := 0, 0
	sum := 0
	for j < k {
		sum += arr[j]
		j++
	}
	avar := sum / k
	if avar >= threshold {
		result++
	}

	lenArr := len(arr)
	for j < lenArr {
		sum -= arr[i]
		i++
		sum += arr[j]
		j++

		avar = sum / k
		if avar >= threshold {
			result++
		}

	}

	return result
}

func main1343() {

}
