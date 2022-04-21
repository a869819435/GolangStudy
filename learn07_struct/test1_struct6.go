package main

import "fmt"

// type 表示申明一个新的数据类型,或者说是int的别名
type Integer int

// 定义一个结构体（类似java的类，与C/C++的结构体概念一致）
type Book struct {
	title string
	auth  string
}

// 这里是一个值传递，无法修改传入的book内容
func changeBook(book Book) {
	book.auth = "1111"
}

// 这里是一个值传递，传递的值为一个指针，里面是原book的地址
func changeBook2(book *Book) {
	book.auth = "1111"
}

func main() {
	var a Integer = 10
	fmt.Println("a = ", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "666"

	fmt.Printf("%v\n", book1)

	changeBook(book1)
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
