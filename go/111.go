package main

import (
	"container/list"
	"fmt"
	"gotest/tree"
	"math"
)

/*
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：2

示例 2：
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5

提示：
树中节点数的范围在 [0, 105] 内
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
func minDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	deque := list.New()
	res := 0
	deque.PushBack(root)

	for deque.Len() > 0 {
		level := deque.Len()
		for i := 0; i < level; i++ {
			node := deque.Front().Value.(*tree.TreeNode)
			deque.Remove(deque.Front())
			if node.Left == nil && node.Right == nil {
				return res + 1
			} else {
				if node.Left != nil {
					deque.PushBack(node.Left)
				}
				if node.Right != nil {
					deque.PushBack(node.Right)
				}
			}
		}
		res++
	}
	return res + 1
}

func main111() {
	a := []int{2, 3, 3, 4, 5, 5, 4, math.MinInt, math.MinInt, 8, 9, 9, 8}
	root := tree.BuildTree(a)
	fmt.Println(minDepth(root))
}
