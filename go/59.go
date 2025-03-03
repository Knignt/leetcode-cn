package main

import (
	"fmt"
)

/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
示例 1：
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
示例 2：
输入：n = 1
输出：[[1]]
*/

func generateMatrix(n int) [][]int {
	// 对于0和1的单独处理
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	len := n - 1
	num := 1
	i, j := 0, 0
	for len > 0 {

		for ; j < len; j++ {
			result[i][j] = num
			num++
		}

		for ; i < len; i++ {
			result[i][j] = num
			num++
		}

		for ; j > j-len-1 && j > 0; j-- {
			result[i][j] = num
			num++
		}

		for ; i > i-len-1 && i > 0; i-- {
			result[i][j] = num
			num++
		}

		len--
	}

	//对于是奇数的处理
	if n/2 > 0 {
		result[n/2][n/2] = num
	}

	return result
}

func main() {
	res := generateMatrix(4)
	fmt.Println(res)
}
