package main

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]

提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

进阶：你能尝试使用一趟扫描实现吗？
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
主要思路在于快慢指针，快指针和慢指针之间的差距为n
需要注意的点是 由于引入了 dammyHead ，所以下标从0开始，并且需要将head也指向dammyHead
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}

	dammyHead := &ListNode{Val: -1, Next: nil}
	dammyHead.Next = head
	fast := dammyHead
	head = dammyHead //head也需要指向到dammyHead
	i := 0
	for i < n {
		fast = fast.Next
		i++
	}

	for fast.Next != nil { //一直挪到最后一个节点的位置
		fast = fast.Next
		head = head.Next
	}

	head.Next = head.Next.Next

	return dammyHead.Next
}

func main19() {
	a := []int{1, 2, 3}
	//a := []int{1, 2}
	b := createList(a)
	c := removeNthFromEnd(b, 2)
	printList(c)
}
