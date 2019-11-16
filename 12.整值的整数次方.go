package main

import "fmt"

//给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。
//保证base和exponent不同时为0
//
func Power(base float64, exponent int) float64 {
	r, p := 1.0, 0
	if exponent < 0 {
		p = -exponent
	} else {
		p = exponent
	}
	for p != 0 {
		//if p & 1 == 1 {
		//	r = r*base
		//}
		//base = base*base
		//p >>= 1
		r = r * base
		p--
	}
	if exponent > 0 {
		return r
	} else {
		return 1 / r
	}
}

func main() {
	p := 3
	b := p & 1
	fmt.Println(b)
}
