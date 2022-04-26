package main

import "fmt"

func main() {
	var a string
	// a 内部就是一个pair<type:string, value:"aaaa">
	a = "aaaa"
	fmt.Println(a)
	var allType interface{}
	allType = a
	// _是占位符，表示不需要获取该变量
	_, str := allType.(string)
	fmt.Println(str)
}
