package main

import "fmt"

/*
给你单链表的头结点 head ，请你找出并返回链表的中间结点。
如果有两个中间结点，则返回第二个中间结点。

示例 1：
输入：head = [1,2,3,4,5]
输出：[3,4,5]
解释：链表只有一个中间结点，值为 3 。

示例 2：
输入：head = [1,2,3,4,5,6]
输出：[4,5,6]
解释：该链表有两个中间结点，值分别为 3 和 4 ，返回第二个结点。

提示：
链表的结点数范围是 [1, 100]
1 <= Node.val <= 100
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Next

}

func main() {

	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(middleNode(createList(a)))
}
