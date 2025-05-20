package main

import "fmt"

/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。



示例 1：


输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]
示例 2：


输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]


提示：

m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1


进阶：

一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个仅使用常量空间的解决方案吗？
*/

/*
解题思路：
先把每一行有0的都扫描出来，然后该行置为0，然后把列号放入一个数组里面，再把数组没置0的列给出来，最后再将列置为0即可.
*/

func setZeroes(matrix [][]int) {
	if matrix == nil {
		return
	}

	row, col := len(matrix), len(matrix[0])

	cols := make([]int, 0)
	for i := 0; i < row; i++ {
		hashZero := false
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				//这一行存在0，直接将这一行全部置为0,然后给出一个[]int，将列重置一下即可
				hashZero = true
				cols = append(cols, j)
				//break			这里不应该break，如果某一行有两个0会有问题
			}
		}

		if hashZero {
			for n := 0; n < col; n++ {
				matrix[i][n] = 0
			}
		}
	}

	fmt.Println(cols)

	if len(cols) > 0 {
		for times := 0; times < len(cols); times++ {
			for i := 0; i < row; i++ {
				matrix[i][cols[times]] = 0
			}
		}
	}
}

func main() {

	a := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	a = [][]int{
		{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5},
	}

	setZeroes(a)

}
