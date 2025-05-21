package main

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。



示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]

示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]


提示：
每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	ans := &ListNode{
		Val:  -1,
		Next: nil,
	}
	temp := ans
	left := 0

	//这里不需要单独考虑某一个链表结束,要是结束了，那么里面对于已结束的节点不做处理,仅仅只关注没结束的
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		s := n1 + n2 + left
		if s >= 0 {
			left = s / 10
			s = s % 10
		}
		node := &ListNode{Val: s}
		temp.Next = node
		temp = temp.Next
	}
	if left == 1 {
		temp.Next = &ListNode{Val: left, Next: nil}
	}

	return ans.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	ans := &ListNode{
		Val:  -1,
		Next: nil,
	}
	tmp := ans
	left := 0

	for l1 != nil && l2 != nil {
		temp := &ListNode{
			Val:  -1,
			Next: nil,
		}

		temp.Val = l1.Val + l2.Val + left
		if temp.Val >= 10 {
			left = temp.Val / 10 //往上进位
			temp.Val -= 10
		} else {
			left = 0
		}

		tmp.Next = temp
		l1 = l1.Next
		l2 = l2.Next
		tmp = tmp.Next
	}

	//判断是否还有剩余的进位
	if l1 != nil {
		leftl1 := l1
		if left != 0 {
			for leftl1.Next != nil {
				temp := leftl1.Val + left
				if temp >= 10 {
					left = temp / 10 //往上进位
					leftl1.Val = temp - 10
				} else {
					left = 0
					leftl1.Val = temp
				}
				leftl1 = leftl1.Next
			}

			if left != 0 {
				temp := left + leftl1.Val
				if temp >= 10 {
					leftl1.Val = temp - 10
					left = temp / 10
					leftl1.Next = &ListNode{
						Val:  left,
						Next: nil,
					}
				} else {
					leftl1.Val = temp
				}

			}
		}
		tmp.Next = l1
	}

	if l2 != nil {

		leftl := l2
		if left != 0 {
			for leftl.Next != nil {
				temp := leftl.Val + left
				if temp >= 10 {
					left = temp / 10 //往上进位
					leftl.Val = temp - 10
				} else {
					left = 0
					leftl.Val = temp
				}
				leftl = leftl.Next
			}

			if left != 0 {
				temp := left + leftl.Val
				if temp >= 10 {
					leftl.Val = temp - 10
					left = temp / 10
					leftl.Next = &ListNode{
						Val:  left,
						Next: nil,
					}
				} else {
					leftl.Val = temp
				}

			}
		}
		tmp.Next = l2
	}

	if l1 == nil && l2 == nil && left != 0 {
		temp := ans
		for temp.Next != nil {
			temp = temp.Next
		}

		temp.Next = &ListNode{
			Val:  left,
			Next: nil,
		}
	}

	return ans.Next
}

func main() {
	a := []int{9, 9, 9, 9, 9, 9, 9}
	b := []int{9, 9, 9, 9}

	a = []int{9, 9, 1}
	b = []int{1}

	a = []int{5}
	b = []int{5}

	a1 := createList(a)
	b1 := createList(b)
	l := addTwoNumbers(a1, b1)
	printList(l)
}
