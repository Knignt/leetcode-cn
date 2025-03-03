package main

import (
	"fmt"
	"gotest/tree"
	"math"
)

/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。



示例 1：


输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]


提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100


进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)

	var stack []tree.TreeNode

	stack = append(stack, *root)
	cur := root

	//中序遍历，是 左 中 右
	//先把所有的左子树push进去，再一个一个弹出时，将右子树push进去
	//再去找右子树的左子树
	//循环往复
	for len(stack) > 0 {
		if cur != nil {
			stack = append(stack, *cur)
			cur = cur.Left
		} else {
			cur = &stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, cur.Val)
			cur = cur.Right
		}
	}

	result = result[:len(result)-1]

	/*var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		//将当前值push进去

		preorder(root.Left)
		result = append(result, root.Val)
		preorder(root.Right)
	}

	preorder(root)*/

	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5, math.MinInt, 8, math.MinInt, math.MinInt, 6, 7, 9}
	root := tree.BuildTree(a)
	fmt.Println(inorderTraversal(root))
}
