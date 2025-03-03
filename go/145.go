package main

import "gotest/tree"

/*
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。



示例 1：

输入：root = [1,null,2,3]

输出：[3,2,1]

解释：



示例 2：

输入：root = [1,2,3,4,5,null,8,null,null,6,7,9]

输出：[4,6,7,5,2,9,8,3,1]

解释：



示例 3：

输入：root = []

输出：[]

示例 4：

输入：root = [1]

输出：[1]



提示：

树中节点的数目在范围 [0, 100] 内
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
func postorderTraversal(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	var stack []tree.TreeNode
	stack = append(stack, *root)

	for len(stack) > 0 {
		tmp := stack[len(stack)-1] //获取栈顶的数据
		stack = stack[:len(stack)-1]
		result = append(result, tmp.Val)

		if tmp.Left != nil { //栈顶的左孩子
			stack = append(stack, *tmp.Left)
		}

		if tmp.Right != nil { //栈顶的右孩子
			stack = append(stack, *tmp.Right)
		}
	}

	s := func(aaa []int) {
		i := 0
		j := len(aaa) - 1

		for i < j {
			aaa[i], aaa[j] = aaa[j], aaa[i]
			i++
			j--
		}
	}

	s(result)

	/*var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		//将当前值push进去
		preorder(root.Left)
		preorder(root.Right)
		result = append(result, root.Val)
	}

	preorder(root)*/

	return result
}

func main145() {

}
