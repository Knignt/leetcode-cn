package main

/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。



示例 1：


输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]
示例 3：

输入：head = [1]
输出：[1]


提示：

链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	result := &ListNode{
		Val:  -1,
		Next: nil,
	}
	result.Next = head
	temp := result
	for head != nil {
		if head.Next == nil {
			break
		}
		tmp := head.Next.Next
		temp.Next = head.Next
		temp.Next.Next = head
		head.Next = tmp
		head = tmp
		temp = temp.Next.Next
	}
	return result.Next
}

func main24() {
	//a := []int{1, 2, 3, 4, 5, 6}
	//a := []int{1}
	a := []int{1, 2, 3}
	b := createList(a)
	c := swapPairs(b)
	printList(c)
}
