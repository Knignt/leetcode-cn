package main

import (
	"fmt"
	"leetcode-cn/tree"
	"math"
)

/*
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

示例 1：
输入：root = [1,null,2,3]
输出：[1,2,3]
解释：

示例 2：
输入：root = [1,2,3,4,5,null,8,null,null,6,7,9]
输出：[1,2,4,5,6,7,3,8,9]
解释：

示例 3：
输入：root = []
输出：[]

示例 4：
输入：root = [1]
输出：[1]

提示：
树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100
*/

func preorderTraversal(root *tree.TreeNode) []int {
	if root == nil {
		return nil
	}

	//迭代方式
	result := make([]int, 0)
	var stack []tree.TreeNode

	stack = append(stack, *root)

	for len(stack) > 0 {
		tmp := stack[len(stack)-1] //获取栈顶的数据
		stack = stack[:len(stack)-1]
		result = append(result, tmp.Val)

		if tmp.Right != nil { //栈顶的右孩子		这样pop出来的话，就是左孩子在前
			stack = append(stack, *tmp.Right)
		}

		if tmp.Left != nil { //栈顶的左孩子
			stack = append(stack, *tmp.Left)
		}
	}

	/*
		递归方式

		var preorder func(root *TreeNode)
		preorder = func(root *TreeNode) {
			if root == nil {
				return
			}
			//将当前值push进去
			result = append(result, root.Val)
			preorder(root.Left)
			preorder(root.Right)
		}

		preorder(root)

	*/

	return result
}

func main144() {
	a := []int{1, 2, 3, 4, 5, math.MinInt, 8, math.MinInt, math.MinInt, 6, 7, 9}
	root := tree.BuildTree(a)
	fmt.Println(preorderTraversal(root))
}
