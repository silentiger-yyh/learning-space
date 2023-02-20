package main

import "fmt"

func main() {
	var a = "string"
	fmt.Println(a)
	b := "sss"
	fmt.Println(b)

	c := a + b
	fmt.Println(c)

	// 常量，没有确定的类型，会根据使用的上下文自动确定类型
	const d = "const"
	const e string = "const str"
	fmt.Println(d)
	fmt.Println(e)
}
