package main

import "fmt"

// 万能类型(空接口) interface{}
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called ...")
	fmt.Println(arg)
	// 判断万能类型的具体类型
	value, ok := arg.(string)
	if ok {
		fmt.Printf("arg is string type,type = %T\n", value)
	}
}

type AA struct {
	str string
}

func main() {
	aa := AA{"???"}

	myFunc(aa)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14159)
}
