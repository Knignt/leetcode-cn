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
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。

示例 1：
输入：root = [3,1,4,null,2], k = 1
输出：1

示例 2：
输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3

提示：
树中的节点数为 n 。
1 <= k <= n <= 104
0 <= Node.val <= 104

进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
*/

func kthSmallest(root *tree.TreeNode, k int) int {
	if root == nil || k <= 0 {
		return k
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	stack := make([]*tree.TreeNode, 0)
	//stack = append(stack, root)
	cur := root
	i := 1

	for cur != nil || len(stack) > 0 {

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == i {
			return cur.Val
		}
		cur = cur.Right
		i++
	}

	return 0

}

func main() {

	a := []int{5, 3, 6, 2, 4, math.MinInt, math.MinInt, 1}
	root := tree.BuildTree(a)
	k := 3
	fmt.Println(kthSmallest(root, k))

}
