package main

import "fmt"

/*
几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。
每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。
你的点数就是你拿到手中的所有卡牌的点数之和。
给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。

示例 1：
输入：cardPoints = [1,2,3,4,5,6,1], k = 3
输出：12
解释：第一次行动，不管拿哪张牌，你的点数总是 1 。但是，先拿最右边的卡牌将会最大化你的可获得点数。最优策略是拿右边的三张牌，最终点数为 1 + 6 + 5 = 12 。

示例 2：
输入：cardPoints = [2,2,2], k = 2
输出：4
解释：无论你拿起哪两张卡牌，可获得的点数总是 4 。

示例 3：
输入：cardPoints = [9,7,7,9,7,7,9], k = 7
输出：55
解释：你必须拿起所有卡牌，可以获得的点数为所有卡牌的点数之和。

示例 4：
输入：cardPoints = [1,1000,1], k = 1
输出：1
解释：你无法拿到中间那张卡牌，所以可以获得的最大点数为 1 。

示例 5：
输入：cardPoints = [1,79,80,1,1,1,200,1], k = 3
输出：202

提示：
1 <= cardPoints.length <= 10^5
1 <= cardPoints[i] <= 10^4
1 <= k <= cardPoints.length

这道题的难处在于题意理解
需要一个滑动窗口从左右，可选择的范围在 左边k开始递减到右边的倒数第k个数，不包括中间的数字，需要注意边界值的处理。
*/
func maxScore(cardPoints []int, k int) int {
	if len(cardPoints) == 0 || k <= 0 {
		return 0
	}

	left, result := 0, 0
	i := 0
	for ; i < k; i++ {
		left += cardPoints[i]
	}
	result = left
	if len(cardPoints) == k {
		return result
	}
	i--

	//从左边减去右边
	j := len(cardPoints) - 1
	for i >= 0 && j >= len(cardPoints)-1-k { //这里要等于，即包含0以及倒数第k个数。
		left -= cardPoints[i]
		left += cardPoints[j]
		i--
		j--
		result = max(result, left)
	}

	return result
}

func main() {
	a := []int{1, 79, 80, 1, 1, 1, 200, 1}
	//a = []int{1, 2, 3, 4, 5, 6, 1}
	k := 3
	fmt.Println(maxScore(a, k))
}
