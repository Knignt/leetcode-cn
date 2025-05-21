package main

import (
	"fmt"
	"math"
)

/*
给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

示例 1:
输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

示例 2:
输入: nums = [2,3,0,1,4]
输出: 2

提示:
1 <= nums.length <= 104
0 <= nums[i] <= 1000
题目保证可以到达 nums[n-1]
*/

func jump(nums []int) int {

	n := len(nums)
	if n <= 1 {
		return 0
	}

	/*
			确定dp数组（dp table）以及下标的含义
			确定递推公式
			dp数组如何初始化
			确定遍历顺序
			举例推导dp数组

		dp 存储到当前这一步 所走的最少得步数

		递推公式：
		dp[j] = min(dp[j], dp[i] + 1)
		如果之前已经走到了这一步，查看当前将要走的这步 和 以前已走的这一步哪一步更小，来进行赋值。
	*/

	dp := make([]int, n)
	for i := range n {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		maxJump := i + nums[i]
		for j := i + 1; j <= maxJump && j < n; j++ {
			dp[j] = min(dp[j], dp[i]+1)
		}
		if maxJump >= n-1 {
			return dp[n-1]
		}
	}
	return dp[n-1]
}

func main() {
	a := []int{2, 3, 1, 1, 4}                              //2
	a = []int{0}                                           //0
	a = []int{2, 3, 0, 1, 4}                               //2
	a = []int{1}                                           //0
	a = []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3} //2
	fmt.Println(jump(a))
}
