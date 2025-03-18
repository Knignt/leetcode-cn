package main

import (
	"fmt"
	"leetcode-cn/tree"
	"math"
)

/*
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，
最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

示例 1:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。

示例 2:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。

说明:
所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉搜索树中。
*/

func lowestCommonAncestor(root, p, q *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}

	//一个比root大，一个比root小，那么就直接返回root
	if (p.Val <= root.Val && q.Val >= root.Val) || (p.Val >= root.Val && q.Val <= root.Val) {
		return root
	}

	//都在左子树或者右子树
	if p.Val < root.Val && q.Val < root.Val {
		root = root.Left
	} else {
		root = root.Right
	}

	stack := make([]*tree.TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		level := len(stack)
		tmp := make([]*tree.TreeNode, 0)
		for i := 0; i < level; i++ {
			if (p.Val <= stack[i].Val && q.Val >= stack[i].Val) || (p.Val >= stack[i].Val && q.Val <= stack[i].Val) {
				return stack[i]
			}

			if stack[i].Left != nil {
				tmp = append(tmp, stack[i].Left)
			}

			if stack[i].Right != nil {
				tmp = append(tmp, stack[i].Right)
			}
		}
		stack = make([]*tree.TreeNode, 0)
		stack = append(stack, tmp...)
	}

	return nil
}

func main235() {
	a := []int{6, 2, 8, 0, 4, 7, 9, math.MinInt, math.MinInt, 3, 5}
	root := tree.BuildTree(a)
	p := &tree.TreeNode{Val: 2}
	q := &tree.TreeNode{Val: 4}
	fmt.Println(lowestCommonAncestor(root, p, q))

}
