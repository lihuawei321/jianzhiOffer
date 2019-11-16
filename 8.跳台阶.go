package main

import "fmt"

//一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

//方法1：递归
func jumpFloor(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	} else {
		return jumpFloor(n-1) + jumpFloor(n-2)
	}

}

//方法2：非递归
func jumpFloor2(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	a, b := 1, 2
	for ; n >= 3; n-- {
		a, b = b, a+b
	}

	return b
}
func main() {
	floor := jumpFloor2(5)
	i := jumpFloor(5)
	fmt.Println(floor)
	fmt.Println(i)
}
