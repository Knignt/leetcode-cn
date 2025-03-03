package main

import (
	"container/list"
	"fmt"
	"leetcode-cn/tree"
	"math"
)

/*
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例 1：
输入：root = [1,2,3,null,5,null,4]
输出：[1,3,4]

示例 2：
输入：root = [1,2,3,4,null,null,null,5]
输出：[1,3,4,5]

示例 3：
输入：root = [1,null,3]
输出：[1,3]

示例 4：
输入：root = []
输出：[]

提示:
二叉树的节点个数的范围是 [0,100]
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
func rightSideView(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	deque := list.New()
	deque.PushBack(root)

	for deque.Len() > 0 {
		level := deque.Len()
		for i := 0; i < level; i++ {
			node := deque.Remove(deque.Front()).(*tree.TreeNode)

			if i == level-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				deque.PushBack(node.Left)
			}
			if node.Right != nil {
				deque.PushBack(node.Right)
			}

		}
	}

	return result
}

func main199() {
	a := []int{1, 2, 3, math.MinInt, 5, math.MinInt, 4}
	root := tree.BuildTree(a)
	fmt.Println(rightSideView(root))
}
