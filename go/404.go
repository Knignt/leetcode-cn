package main

import (
	"fmt"
	"leetcode-cn/tree"
	"math"
)

/*
给定二叉树的根节点 root ，返回所有左叶子之和。

示例 1：
输入: root = [3,9,20,null,null,15,7]
输出: 24
解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

示例 2:
输入: root = [1]
输出: 0

提示:
节点数在 [1, 1000] 范围内
-1000 <= Node.val <= 1000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}

	result := make([]int, 0)
	var stack []*tree.TreeNode
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			stack = append(stack, node.Left)
			left := node.Left
			if left != nil && left.Left == nil && left.Right == nil {
				result = append(result, left.Val)
			}
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	s := 0
	for i := 0; i < len(result); i++ {
		s += result[i]
	}
	return s
}

func main404() {

	a := []int{1, 2, 3, math.MinInt, 5}
	root := tree.BuildTree(a)
	fmt.Println(sumOfLeftLeaves(root))

}
