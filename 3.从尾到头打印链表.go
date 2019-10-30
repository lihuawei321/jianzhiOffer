package main

//输入一个链表的头结点，按照 从尾到头 的顺序返回节点的值。
//返回的结果用数组存储。
type ListNode struct {
	Val  int
	Next *ListNode
}

func printListReversingly(head *ListNode) []int {
	var res []int
	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return reverse(res)
}

func reverse(res []int) []int {
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res

}

func main() {

}
