package main

import (
	"fmt"
	"strings"
)

//请实现一个函数，将一个字符串中的每个空格替换成“%20”。
// 例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。

func replaceSpaces(str string) string {
	var res []string
	for _, v := range str {
		if v == ' ' {
			res = append(res, "%20")
		} else {
			res = append(res, string(v))
		}
	}
	return strings.Join(res,"")
}
func main() {
	s:="We are happy."
	spaces := replaceSpaces(s)
	fmt.Println(spaces)
}
