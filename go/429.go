package main

/*
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

示例 1：
输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]

示例 2：
输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]

提示：
树的高度不会超过 1000
树的节点总数在 [0, 104] 之间
*/

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	var result [][]int
	list := make([]*Node, 0)
	list = append(list, root)

	tmp := make([]*Node, 0)
	path := make([]int, 0)
	for len(list) > 0 {
		node := list[len(list)-1]
		list = list[1:]
		if node.Children != nil {
			for i := 0; i < len(node.Children); i++ {
				tmp = append(tmp, node.Children[i])
			}
		}
		if len(list) > 0 {
			path = append(path, node.Val)
		} else {
			list = tmp
			result = append(result, path)
			tmp = make([]*Node, 0)
			path = make([]int, 0)
		}
	}

	return result
}

func main() {

}
