package main

import "fmt"

// 关键字 函数名(形参1 类型1,形参2 类型2,...) (返回值类型1，返回值类型2) {
//      函数体[具体函数功能编写]
// }
func add(a int, b int) int {
	return a + b
}
// 返多个匿名返回值
func addAndSub(a int, b int) (int, int) {
	return a + b, a - b
}

// 返多个非匿名返回值
func addAndSub2(a int, b int) (ans1 int, ans2 int) {
	// 未赋值之前，非匿名为局部变量，初始值为0，可以看做函数的形参
	ans1 = add(a,b)
	ans2 = a - b
	return ans1, ans2
}

// 返多个非匿名返回值,同类型返回参数可以放一起写
func addAndSub3(a int, b int) (ans1, ans2 int, str string) {
	ans1 = add(a,b)
	ans2 = a - b
	return ans1, ans2, "123"
}

func main()  {
	c := add(5,6)
	fmt.Println(c)
	r1, r2 := addAndSub(8,6)
	fmt.Println(r1,",",r2)
	r3, r4 := addAndSub2(8,6)
	fmt.Println(r3,",",r4)
}
