package main

import "fmt"

func pr4(array [4]int)  {
	fmt.Println("ok4")
	array[0] = 1000
}

func pr10(array [10]int)  {
	fmt.Println("ok10")
}

func prDiy(array []int)  {
	fmt.Println("ok diy")
	array[0] = 817
}

func main()  {
	// 固定长度的数组,默认数组内的值都为0
	var array1 [10]int
	for i := 0; i < len(array1) ; i++ {
		fmt.Println(array1[i])
	}

	// 固定长度的数组,初始化前四个值
	array2 := [10]int{1,2,3,4}
	// range 关键字会根据选中的不同的j集合去访问不同的值
	for i, value := range array2 {
		fmt.Println("index = ", i, ",value =" ,value)
	}

	// 参数传递的过程中，数组的大小一定要匹配，否则编译不通过
	array3 := [4]int{11,22,33,44}
	for i := range array3 {
		fmt.Println(i)
	}
	fmt.Println(array3[0]) // 结果为11
	pr4(array3)
	fmt.Println(array3[0])  // 结果仍未为 11
	// pr10(array3)

	// 切片数组
	array4 := []int{1,2,3,4}
	fmt.Printf("array4 type of %T\n", array4)
	// 切片数组的形参传递的是地址，即引用传递，因此可以通过数组
	// 去改变对应的数组值[动态数据本身就是一个指向这个块空间的指针]
	fmt.Println(array4[0])
	prDiy(array4)
	// pr10(array4)
	// prDiy(array3)
	fmt.Println(array4[0])
}
