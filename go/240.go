package main

import (
	"fmt"
)

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
每行的元素从左到右升序排列。
每列的元素从上到下升序排列。


示例 1：
输入：matrix = [
[1,4,7,11,15],
[2,5,8,12,19],
[3,6,9,16,22],
[10,13,14,17,24],
[18,21,23,26,30]
], target = 5
输出：true

示例 2：
输入：matrix = [
[1,4,7,11,15],
[2,5,8,12,19],
[3,6,9,16,22],
[10,13,14,17,24],
[18,21,23,26,30]
], target = 20
输出：false

提示：
m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-109 <= matrix[i][j] <= 109
每行的所有元素从左到右升序排列
每列的所有元素从上到下升序排列
-109 <= target <= 109
*/

/*
从右上角处理,如果当前值等于直接回true
如果小于，则说明当前列还可以再++
如果大于，则说明当前列不可能会存在比target小的值了，主要因为是从右上角开始的

为什么选择右下角？
如果从左上角或者左下角开始，没办法知道轮训的方向
1 2 6
3 7 8
如果找3,从左上角，找到的就是2这一行了而不是1这一行
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || matrix == nil {
		return false
	}

	row, col := len(matrix), len(matrix[0])
	i, j := 0, col-1
	for i < row && j >= 0 {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}

	return false
}

func main() {
	a := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	/*a = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	a = [][]int{
		{-5},
	}*/

	target := -5
	//target = 20
	//target = -2
	//target = -5
	target = -10
	target = 19
	fmt.Println(searchMatrix(a, target))
}
