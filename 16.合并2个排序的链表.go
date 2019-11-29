package main

//输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。

//递归
func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge(l1, l2.Next)
		return l2
	}
	return nil

}

//非递归
func merge1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	var cur *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			if head == nil {
				cur = l1
				head = cur
			} else {
				cur.Next = l1
				cur = cur.Next
			}
			l1 = l1.Next
		} else {
			if head == nil {
				cur = l2
				head = cur
			} else {
				cur.Next = l2
				cur = cur.Next
			}
			l2 = l2.Next
		}
		if l1 == nil {
			cur.Next = l2
		} else {
			cur.Next = l1
		}
	}
	return head
}
func main() {

}
