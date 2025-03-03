package main

import (
	"container/list"
	"fmt"
	"leetcode-cn/tree"
	"math"
)

/*
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
假设二叉树中至少有一个节点。

示例 1:
输入: root = [2,1,3]
输出: 1

示例 2:
输入: [1,2,3,4,null,5,6,null,null,7]
输出: 7

提示:
二叉树的节点个数的范围是 [1,104]
-231 <= Node.val <= 231 - 1
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
层序遍历获取每一层第一个值
*/
func findBottomLeftValue(root *tree.TreeNode) int {
	if root == nil {
		return -1
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	queue := list.New()
	queue.PushBack(root)

	result := -1
	for queue.Len() > 0 {
		level := queue.Len()
		tmp := list.New()
		for i := 1; i <= level; i++ {
			node := queue.Front().Value.(*tree.TreeNode)
			queue.Remove(queue.Front())

			if node.Left != nil {
				tmp.PushBack(node.Left)
			}
			if node.Right != nil {
				tmp.PushBack(node.Right)
			}
			if i == 1 {
				result = node.Val
			}
		}
		queue = tmp
	}

	return result
}

func main513() {
	a := []int{0, math.MinInt, -1} //-1
	//a = []int{2, 1, 3}             //1
	//a = []int{1, 2, 3, 4, math.MinInt, 5, 6, math.MinInt, math.MinInt, 7} //7a
	a = []int{3, 1, 5, 0, 2, 4, 6} //0
	root := tree.BuildTree(a)
	fmt.Println(findBottomLeftValue(root))
}
