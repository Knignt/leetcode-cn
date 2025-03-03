package main

import (
	"fmt"
	"leetcode-cn/tree"
	"math"
	"strconv"
)

/*
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。


示例 1：
输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]

示例 2：
输入：root = [1]
输出：["1"]

提示：
树中节点的数目在范围 [1, 100] 内
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

func binaryTreePaths(root *tree.TreeNode) []string {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	var stack []*tree.TreeNode
	stack = append(stack, root)
	var result []string
	paths := []string{strconv.Itoa(root.Val)} //记录当前已遍历的节点

	/*
		每一个栈顶的节点pop出来，再将其右左孩子push进去，之所以是右左孩子，是因为这样的话，输出顺序是先左边再右边
		stack 保存待遍历的点
		paths 保存前面已遍历的节点路径
		这里不需要担心没对上，主要是因为push时是一起push进去的
	*/

	for len(stack) > 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		path := paths[len(paths)-1]
		paths = paths[:len(paths)-1]

		if node.Left == nil && node.Right == nil {
			result = append(result, path)
			continue
		}

		//为什么这里要这么写呢？要保证stack和paths里面的值遍历顺序一致
		//即stack遍历过该节点,也需要放到paths里面去
		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+"->"+strconv.Itoa(node.Right.Val))
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
			paths = append(paths, path+"->"+strconv.Itoa(node.Left.Val))
		}

	}
	return result
}

func main257() {
	a := []int{1, 2, 3, math.MinInt, 5}
	root := tree.BuildTree(a)
	fmt.Println(binaryTreePaths(root))
}
