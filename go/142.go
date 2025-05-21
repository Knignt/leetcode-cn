package main

/*
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。


示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。

提示：
链表中节点的数目范围在范围 [0, 104] 内
-105 <= Node.val <= 105
pos 的值为 -1 或者链表中的一个有效索引
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
假设慢指针走了x 步，那么快指针走了2x步，那么 2x-x = nb (这里b是环的长度，n是快指针多走的n圈)，
x= nb，即慢指针走了 n圈的环的长度，但是这里还包含了一个未在环的长度a，即慢指针在环里走了nb-a步，
要使慢指针走的刚好是环的倍数的长度，需要往前走a步，这样刚好是n个环的长度，也就刚好走到了环的入口处，
即我们只需要让head和 slow都走a步，就刚好在入口处，反过来就是说 head和slow 相等的地方就是环的入口处。
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { //这里只是相遇点而已，不是环的入口
			for head != slow {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}

func main() {

}
