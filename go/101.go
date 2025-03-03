package main

import (
	"container/list"
	"fmt"
	"leetcode-cn/tree"
	"math"
)

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。

示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：
输入：root = [1,2,2,null,3,null,3]
输出：false


提示：
树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100
进阶：你可以运用递归和迭代两种方法解决这个问题吗？
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
解题思路：
从第二层开始遍历，从队列前后对比，如果不相等直接false
每遍历两个节点，将其左右孩子对称的形式Push到temp的队列中去
deque 比较完毕，将temp直接赋值给deque
循环结束的结果是 deque比较完毕即可
*/

func isSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return false
	}

	deque := list.New()
	deque.PushFront(root.Right)
	deque.PushFront(root.Left)

	for deque.Len() > 0 {

		level := deque.Len()
		j := level - 1
		tmp := list.New()

		for i := 0; i < j; i++ {
			if deque.Front().Value.(*tree.TreeNode) == nil && deque.Back().Value.(*tree.TreeNode) == nil {
				deque.Remove(deque.Front())
				deque.Remove(deque.Back())
				j--
				continue
			}

			if (deque.Front() == nil && deque.Back() != nil) || (deque.Front() != nil && deque.Back() == nil) || (deque.Front().Value.(*tree.TreeNode) != nil && deque.Back().Value.(*tree.TreeNode) == nil) || (deque.Front().Value.(*tree.TreeNode) == nil && deque.Back().Value.(*tree.TreeNode) != nil) {
				return false
			}

			if deque.Front().Value.(*tree.TreeNode).Val != deque.Back().Value.(*tree.TreeNode).Val {
				return false
			} else {
				//push的顺序要修改，这样子left是在前的，
				tmp.PushFront(deque.Front().Value.(*tree.TreeNode).Right)
				tmp.PushFront(deque.Front().Value.(*tree.TreeNode).Left)

				tmp.PushBack(deque.Back().Value.(*tree.TreeNode).Left)
				tmp.PushBack(deque.Back().Value.(*tree.TreeNode).Right)
				deque.Remove(deque.Front())
				deque.Remove(deque.Back())
				j--
			}
		}
		deque = tmp

	}

	return true
}

func main101() {
	//a := []int{1, 2, 2, 3, 4, 4, 3}
	//a := []int{1, 2, 2, math.MinInt, 3, math.MinInt, 3}
	//a := []int{1, 2, 2, math.MinInt, 3, 3}
	a := []int{2, 3, 3, 4, 5, 5, 4, math.MinInt, math.MinInt, 8, 9, 9, 8}
	root := tree.BuildTree(a)
	fmt.Println(isSymmetric(root))
}
