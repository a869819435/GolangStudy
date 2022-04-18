package main

import "fmt"

func main()  {
	// 声明一个map，key是string, value是int类型
	var map1 map[string]int
	if map1 == nil {
		fmt.Println("空的map")
	}
	// 声明一个带有空间的map
	map2 := make(map[string]int,10)
	if map2 != nil {
		fmt.Println("非空的map")
	}

	// 赋值方式
	map2["a"] = 1
	map2["b"] = 2
	map2["c"] = 3
	fmt.Println(map2)

	// 声明一个开空间且有元素值的map
	map3 := map[string]int {
		"a" : 5,
		"b" : 6,
		"c" : 7,
	}
	fmt.Println(map3)

}
