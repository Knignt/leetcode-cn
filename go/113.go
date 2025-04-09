package main

import (
	"fmt"
	"leetcode-cn/tree"
	"math"
	"os"
)

/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。



示例 1：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]

示例 2：
输入：root = [1,2,3], targetSum = 5
输出：[]

示例 3：
输入：root = [1,2], targetSum = 0
输出：[]

提示：
树中节点总数在范围 [0, 5000] 内
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000
*/

func pathSum(root *tree.TreeNode, targetSum int) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	result := make([][]int, 0)

	path := make([]*tree.TreeNode, 0)
	stack := make([]*tree.TreeNode, 0)
	sumStack := make([]int, 0)

	stack = append(stack, root)
	path = append(path, root)
	sumStack = append(sumStack, root.Val)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nodeSum := sumStack[len(sumStack)-1]
		sumStack = sumStack[:len(sumStack)-1]

		path = path[:len(path)-1]

		if nodeSum+node.Val == targetSum && node.Left == nil && node.Right == nil {
			tmp := make([]int, 0)
			for i := 0; i < len(path); i++ {
				tmp = append(tmp, path[i].Val)
			}
			tmp = append(tmp, node.Val)
			result = append(result, tmp)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
			sumStack = append(sumStack, nodeSum+node.Left.Val)
			path = append(path, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			sumStack = append(sumStack, nodeSum+node.Right.Val)
			path = append(path, node.Right)
		}

		fmt.Println(stack, path, sumStack)
		os.Exit(1)

	}

	return result
}

func main() {

	a := []int{5, 4, 8, 11, math.MinInt, 13, 4, 7, 2, math.MinInt, math.MinInt, 5, 1}
	root := tree.BuildTree(a)
	targetSum := 22
	fmt.Println(pathSum(root, targetSum))

}
