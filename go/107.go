package main

import (
	"container/list"
	"fmt"
	"gotest/tree"
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

func levelOrderBottom(root *tree.TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	deque := list.New()
	deque.PushBack(root)

	for deque.Len() > 0 {
		level := deque.Len()

		tmp := make([]int, 0)
		for i := 0; i < level; i++ {
			node := deque.Remove(deque.Front()).(*tree.TreeNode)
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				deque.PushBack(node.Left)
			}
			if node.Right != nil {
				deque.PushBack(node.Right)
			}
		}

		result = append(result, tmp)
	}

	reseve := func(a [][]int) {
		i := 0
		j := len(a) - 1
		for i < j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}

	reseve(result)

	return result
}

func main107() {
	a := []int{3, 9, 20, math.MinInt, math.MinInt, 15, 7}
	root := tree.BuildTree(a)
	fmt.Println(levelOrderBottom(root))
}
