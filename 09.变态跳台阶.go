package main

//一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
//方法1：递归
func btjumpFloor(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	} else {
		return 2 * btjumpFloor(n-1)
	}

}

//方法2：非递归
func btjumpFloor2(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	b := 2
	for ; n >= 3; n-- {
		b = 2 * b
	}

	return b
}
func main() {

}
