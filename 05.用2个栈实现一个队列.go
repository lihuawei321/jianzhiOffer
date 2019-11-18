package main

//用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。

type MyQueue struct {
	InputStack  []int
	OutputStack []int
}

func Constructor() MyQueue {
	queue := MyQueue{[]int{}, []int{}}
	return queue
}

func (this *MyQueue) Push(x int) {
	this.InputStack = append(this.InputStack, x)
}

func (this *MyQueue) Pop() int {
	//当输出栈不为空时，那就直接取出输出栈的栈顶元素
	if len(this.OutputStack) != 0 {
		x := this.OutputStack[len(this.OutputStack)-1]
		this.OutputStack = this.OutputStack[:len(this.OutputStack)-1]
		return x
	} else if len(this.InputStack) != 0 {
		for i := len(this.InputStack); i > 0; i-- {
			this.OutputStack = append(this.OutputStack, this.InputStack[i-1])
			this.InputStack = this.InputStack[:len(this.InputStack)-1]
		}
		x := this.OutputStack[len(this.OutputStack)-1]
		this.OutputStack = this.OutputStack[:len(this.OutputStack)-1]
		return x
	}
	return 0 //2个栈都是空的

}

func (this *MyQueue) Peek() int {
	if len(this.OutputStack) != 0 {
		return this.OutputStack[len(this.OutputStack)-1]
	}
	if len(this.InputStack) != 0 {
		for i := len(this.InputStack); i > 0; i-- {
			this.OutputStack = append(this.OutputStack, this.InputStack[i-1])
			this.InputStack = this.InputStack[:len(this.InputStack)-1]
		}
		return this.OutputStack[len(this.OutputStack)-1]
	}
	return 0

}

func (this *MyQueue) Empty() bool {
	if len(this.OutputStack)+len(this.InputStack) == 0 {
		return true
	}
	return false

}

func main() {

}
