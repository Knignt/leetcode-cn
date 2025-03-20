package main

import (
	"fmt"
	"leetcode-cn/tree"
	"math"
)

/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。


示例 1：


输入：root = [2,1,3]
输出：true
示例 2：


输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。


提示：

树中节点数目范围在[1, 104] 内
-231 <= Node.val <= 231 - 1
*/

/*
中序遍历结果天然具有严格递增的特性，所以选择中序遍历，这样的话，只需要从最左边的叶子节点开始，查看是否是一个递增的序列即可
*/

func isValidBST(root *tree.TreeNode) bool {
	if root == nil {
		return false
	}

	stack := make([]*tree.TreeNode, 0)
	//stack = append(stack, root)		这里放进去会导致后续会被pop出来,导致会被计算两次
	var prev *int
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil && cur.Val <= *prev {
			return false
		}

		prev = &cur.Val
		cur = cur.Right

	}

	return true
}

/*
func isValidBST(root *TreeNode) bool {
    var prev *int // 记录前一个节点的值
    stack := []*TreeNode{}
    curr := root

    for curr != nil || len(stack) > 0 {
        // 左子树全部入栈
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        // 处理当前节点
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // 比较当前节点与前一个节点
        if prev != nil && curr.Val <= *prev {
            return false
        }
        // 更新 prev
        val := curr.Val
        prev = &val

        // 处理右子树
        curr = curr.Right
    }
    return true
}
*/

func main() {
	a := []int{1, 1}                                   //false
	a = []int{0, -1}                                   //true
	a = []int{5, 4, 6, math.MinInt, math.MinInt, 3, 7} //false(右侧3的值比root小)
	root := tree.BuildTree(a)
	fmt.Println(isValidBST(root))
}
