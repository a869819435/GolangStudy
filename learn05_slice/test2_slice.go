package main

import "fmt"

func main()  {
	// 声明切片，并且初始化，默认值1,2,3长度len是3
	slice1 := []int{1,2,3}
	fmt.Printf("len = %d, slice = %v\n",len(slice1), slice1)

	// 声明切片，并且初始长度为0
	var slice2 []int
	// 容量为0
	fmt.Printf("len = %d, slice = %v\n",len(slice2), slice2)
	// 需要申请空, 默认值为0  或者 使用
	slice2 = make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n",len(slice2), slice2)

	// 第一第二种结合
	var slice3 = make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n",len(slice3), slice3)

	// 简写
	slice4 := make([]int, 3)
	if slice4 == nil {
		fmt.Println("空切片")
	} else {
		fmt.Printf("len = %d, slice = %v\n",len(slice4), slice4)
	}
}
