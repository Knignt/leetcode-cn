package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	var i int
	dp[0] = nums[0]
	if nums[1] > nums[0] {
		dp[1] = nums[1]
	} else {
		dp[1] = nums[0]
	}
	for i = 2; i < len(nums); i++ {
		tmp := dp[i-2] + nums[i]
		if dp[i-1] > tmp {
			dp[i] = dp[i-1]
		} else {
			dp[i] = tmp
		}
	}

	return dp[i-1]
}

func main() {
	var nums []int = []int{1, 2, 3, 1}
	res := rob(nums)
	fmt.Println(res)
}
