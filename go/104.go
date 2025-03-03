package main

import (
	"container/list"
	"fmt"
	"gotest/tree"
	"math"
)

/*
给定一个二叉树 root ，返回其最大深度。

二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。



示例 1：





输入：root = [3,9,20,null,null,15,7]
输出：3
示例 2：

输入：root = [1,null,2]
输出：2


提示：

树中节点的数量在 [0, 104] 区间内。
-100 <= Node.val <= 100
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	deque := list.New()
	deque.PushBack(root)

	res := 0
	for deque.Len() > 0 {
		level := deque.Len()
		for i := 0; i < level; i++ {
			node := deque.Front().Value.(*tree.TreeNode)
			deque.Remove(deque.Front())
			if node.Left != nil {
				deque.PushBack(node.Left)
			}
			if node.Right != nil {
				deque.PushBack(node.Right)
			}
		}
		res++
	}
	return res
}

func main104() {
	a := []int{2, 3, 3, 4, 5, 5, 4, math.MinInt, math.MinInt, 8, 9, 9, 8}
	root := tree.BuildTree(a)
	fmt.Println(maxDepth(root))
}
