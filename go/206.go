package main

/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。


示例 1：


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例 2：


输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]


提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{Val: -1, Next: nil}
	result := newHead

	for head != nil {
		temp := head
		head = head.Next
		temp.Next = newHead.Next
		newHead.Next = temp
	}

	return result.Next
}

func main206() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := createList(a)
	//printList(b)
	c := reverseList(b)
	printList(c)
}
