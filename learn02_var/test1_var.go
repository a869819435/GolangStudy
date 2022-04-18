package main

import "fmt"

/*
	四种变量的声明方式
 */
func main() {
	// 方法1：默认值为0
	var a int
	fmt.Print("a = ", a)
	// 获取类型,使用printf
	fmt.Printf(",Type of %T\n", a)

	// 方法2：初始化
	var b string = "1"
	fmt.Print("b = ", b)
	fmt.Printf(",Type of %T\n", b)

	// 方法3：在初始化的时候，可以省去数据类型，根据数据值自动匹配类型（不推荐）
	var c = 3.1415962
	fmt.Print("c = ",c)
	fmt.Printf(",Type of %T\n", c)

	// 方法4：省去var关键字，直接自动匹配，只适合局部变量
	d := "100"
	fmt.Print("d = ",d)
	fmt.Printf(",Type of %T\n", d)

	// 打印全局变量
	fmt.Println("gA = ",gA)
	fmt.Println("gB = ",gB)
	fmt.Println("gC = ",gC)

	// 声明多个变量，同类型
	var xx, yy int = 100, 200
	fmt.Println("xx = ",xx,",yy = ",yy)

	// 声明多变量，不同类型
	var kk, ll = 1.3, "aaa"
	fmt.Println("kk = ",kk,",ll = ",ll)

	// 多行的多变量声明
	var (
		vv int = 1001
		ww bool = true
	)
	fmt.Println("vv = ",vv,",ww = ",ww)


}

// 声明全局变量,方法1-3都可以，第四种不行
var gA int
var gB int = 100
var gC = 3.1415962