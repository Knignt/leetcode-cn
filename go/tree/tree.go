package tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTreeNode(b int) *TreeNode {
	return &TreeNode{
		Val:   b,
		Left:  nil,
		Right: nil,
	}
}

func BuildTree(a []int) *TreeNode {
	root := createTreeNode(a[0])
	result := root
	i := 1

	deque := []*TreeNode{root}

	/*
		deque 里面存储的是当前的根节点
		每次循环之后需要将这个head pop出去
	*/

	for i < len(a) {

		node := deque[0]
		deque = deque[1:] //将最前面的pop出去

		//先给左子树处理
		if i < len(a) && a[i] != math.MinInt {
			left := createTreeNode(a[i])
			node.Left = left
			deque = append(deque, node.Left)
		}
		i++

		//再给右子树处理
		if i < len(a) && a[i] != math.MinInt {
			right := createTreeNode(a[i])
			node.Right = right
			deque = append(deque, node.Right)
		}
		i++
	}

	return result
}
