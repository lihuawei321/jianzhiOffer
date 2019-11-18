package main

//输入一个链表，输出该链表中倒数第k个结点。

//使用两个指针，此方法的思想是在查找过程中，
// 设置两个指针，让其中一个指针比另一个指针先前移 k 步，然后两个指针同时往前移动。
// 循环直到先行的指针值为 null 时，另一个指针的位置就是所要找的位置

func findKthToTail(pListHead *ListNode, k int) *ListNode {
	if pListHead == nil {
		return nil
	}
	p1 := pListHead
	for ; k > 0; k-- {
		if p1 != nil {
			p1 = p1.Next
		} else {
			return nil
		}
	}
	p2 := pListHead
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
func main() {

}
