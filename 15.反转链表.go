package main

//输入一个链表，反转链表后，输出新链表的表头。

//非递归
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

//递归
func reverseList2(head *ListNode) *ListNode {
	var newHead *ListNode
	if head == nil || head.Next == nil {
		return head
	} else {
		newHead = reverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
	}
	return newHead
}
func main() {

}
