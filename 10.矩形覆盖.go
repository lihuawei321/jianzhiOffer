package main

//我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？

func count(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	} else {
		return count(n-1) + count(n-2)
	}

}
func main() {

}
