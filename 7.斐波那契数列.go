package main

//大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0）。
//n<=39

//递归
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

//非递归
func Fibonacci2(n int) int {
	a, b := 0, 1
	for ; n > 0; n-- {
		a, b = b, a+b
	}
	return a
}
func main() {

}
