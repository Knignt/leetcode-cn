package main

import (
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
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。


示例 1：


输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [0]
输出：[0]


提示：

树中结点数在范围 [0, 2000] 内
-100 <= Node.val <= 100


进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
*/

func flatten(root *tree.TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}

	stack := make([]*tree.TreeNode, 0)
	var prev *tree.TreeNode
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil {
			prev.Right = node
			prev.Left = nil
		}

		prev = node
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

/*func flatten(root *tree.TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}

	stack := make([]*tree.TreeNode, 0)
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			fmt.Println(cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
}*/

func main() {

	a := []int{1, 2, 5, 3, 4, math.MinInt, 6}
	root := tree.BuildTree(a)
	flatten(root)

}
