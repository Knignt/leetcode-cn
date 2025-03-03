package main

import (
	"fmt"
	"leetcode-cn/tree"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
给定一个二叉树，判断它是否是 平衡二叉树
平衡二叉树 是指该树所有节点的左右子树的高度相差不超过 1。

示例 1：
输入：root = [3,9,20,math.MinInt,math.MinInt,15,7]
输出：true

示例 2：
输入：root = [1,2,2,3,3,math.MinInt,math.MinInt,4,4]
输出：false

示例 3：
输入：root = []
输出：true

提示：
树中的节点数在范围 [0, 5000] 内
-104 <= Node.val <= 104
*/

func isBalanced(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	left := calTreeHigh(root.Left)
	right := calTreeHigh(root.Right)

	if root.Left == nil {
		right++
	}

	if root.Right == nil {
		left++
	}

	return abs(left, right) <= 1 && left >= 0 && right >= 0
}

func abs(a int, b int) int {
	c := a - b
	if c < 0 {
		return c * -1
	} else {
		return c
	}
}

func calTreeHigh(node *tree.TreeNode) int {
	if node == nil {
		return 0
	}

	result := 0
	var stack []*tree.TreeNode
	stack = append(stack, node)

	for len(stack) > 0 {
		itemNode := stack[len(stack)-1]
		if itemNode == nil {
			break
		}
		stack = stack[:len(stack)-1]

		if itemNode.Left == nil && itemNode.Right == nil {
			return result
		}

		if itemNode.Left != nil {
			stack = append(stack, itemNode.Left)
		}

		if itemNode.Right != nil {
			stack = append(stack, itemNode.Left)
		}

		result++
	}

	return result
}

func main() {
	a := []int{3, 9, 20, math.MinInt, math.MinInt, 15, 7} //true
	//a = []int{1, 2, 2, 3, 3, math.MinInt, math.MinInt, 4, 4}                           //false
	//a = []int{1, 2, 2, 3, math.MinInt, math.MinInt, 3, 4, math.MinInt, math.MinInt, 4} //false
	//a = []int{1, math.MinInt, 2, math.MinInt, 3}
	a = []int{1, 2, 2, 3, math.MinInt, math.MinInt, 3, 4, math.MinInt, math.MinInt, 4}
	root := tree.BuildTree(a)
	fmt.Println(isBalanced(root))
}
