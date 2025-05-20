package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headB == nil || headA == nil {
		return nil
	}

	p, q := headA, headB
	isChange := false

	for p != nil || q != nil {
		if p.Next == nil && !isChange {
			p = headA
			isChange = true
		}

		if q.Next == nil && !isChange {
			p = headB
			isChange = true
		}

		p = p.Next
		q = q.Next
	}

	if p == nil {
		return q
	} else {
		return p
	}

}

func main() {

}
