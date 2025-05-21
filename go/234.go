package main

/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
示例 1：
输入：head = [1,2,2,1]
输出：true

示例 2：
输入：head = [1,2]
输出：false

提示：
链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9

进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
先找到最中间的节点，然后中间节点往后全部翻转，然后head 和mid 都进行值的判断，查看是否值一致.
*/

func middleNode1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	result := &ListNode{
		Val:  -1,
		Next: nil,
	}
	newHead := result

	for head != nil {
		temp := head.Next
		head.Next = newHead.Next
		head.Next = head
		head = temp
	}

	return result.Next
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	mid := middleNode1(head)
	mid = reverseList(mid)

	//head 和mid的对比
	ans := false
	for head != nil && mid != nil {
		if head.Val != mid.Val {
			ans = false
			break
		} else {
			ans = true
		}
		head = head.Next
		mid = mid.Next
	}

	return ans
}

func main() {

}
