package main

import "fmt"

func func1()  {
	fmt.Println("A")
}

func main()  {
	// defer但前函数执行完之后必执行的部分
	// 顺序入栈,调用函数也是一样的
	defer fmt.Println("end1")
	defer func1()
	fmt.Println("1523456")

}
