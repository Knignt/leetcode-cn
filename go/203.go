package main

import (
	"fmt"
)

/*
203.移除链表元素
力扣题目链接(opens new window)

题意：删除链表中等于给定值 val 的所有节点。

示例 1： 输入：head = [1,2,6,3,4,5,6], val = 6 输出：[1,2,3,4,5]

示例 2： 输入：head = [], val = 1 输出：[]

示例 3： 输入：head = [7,7,7,7], val = 7 输出：[]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	res := &ListNode{ //给一个虚拟头节点
		Val:  -1,
		Next: nil,
	}
	tmp := res
	for head != nil { //如果值和val不相等,将不相等得连接再最后得结果链表得后面,
		if head.Val != val {
			tmp.Next = head
			tmp = head
			head = head.Next
			tmp.Next = nil
		} else {
			head = head.Next
		}
	}

	return res.Next
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func main203() {
	//先是数组,转成链表
	a := []int{1, 2, 3, 4, 5, 7, 4, 6, 7}
	//a := []int{7, 7, 7, 7, 7}
	//a := []int{1, 1, 1, 1, 2}
	b := createList(a)
	//printList(b)
	c := removeElements(b, 7)
	printList(c)
}

func createList(ints []int) *ListNode {

	head := &ListNode{
		Val:  ints[0],
		Next: nil,
	}
	res := head
	for i := 1; i < len(ints); i++ {
		tmp := &ListNode{
			Val:  ints[i],
			Next: nil,
		}
		head.Next = tmp
		head = head.Next
	}

	return res
}
