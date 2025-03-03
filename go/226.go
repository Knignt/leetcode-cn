package main

import (
	"container/list"
	"gotest/tree"
)

/*
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

示例 1：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

示例 2：
输入：root = [2,1,3]
输出：[2,3,1]

示例 3：
输入：root = []
输出：[]

提示：
树中节点数目范围在 [0, 100] 内
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
/*
翻转二叉树思路：
每一个节点的左右孩子翻转即可
建立一个队列，每一个节点进入到队列中，先翻转该节点的左右孩子
再将左右孩子放入
为什么不会又翻转回去呢？
因为左右孩子已经没关联了
*/
func invertTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}

	deque := list.New()
	deque.PushBack(root)

	for deque.Len() > 0 {
		//交换node的左右孩子，然后push进去
		node := deque.Remove(deque.Front()).(*tree.TreeNode)
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			deque.PushBack(node.Left)
		}
		if node.Right != nil {
			deque.PushBack(node.Right)
		}
	}

	return root
}

func main226() {

}
