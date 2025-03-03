package main

import (
	"container/list"
	"fmt"
	"leetcode-cn/tree"
	"math"
)

/*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]

示例 2：
输入：root = [1]
输出：[[1]]

示例 3：
输入：root = []
输出：[]

提示：

树中节点数目在范围 [0, 2000] 内
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

func levelOrder(root *tree.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	deque := list.New()
	deque.PushFront(root)

	result := make([][]int, 0)

	for deque.Len() > 0 {
		level := deque.Len()
		tmp := make([]int, 0)
		for i := 0; i < level; i++ {
			node := deque.Remove(deque.Front()).(*tree.TreeNode)
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				deque.PushBack(node.Left)
			}
			if node.Right != nil {
				deque.PushBack(node.Right)
			}
		}
		result = append(result, tmp)
	}

	return result
}

func main102() {
	a := []int{3, 9, 20, math.MinInt, math.MinInt, 15, 7}
	root := tree.BuildTree(a)
	fmt.Println(levelOrder(root))
}
