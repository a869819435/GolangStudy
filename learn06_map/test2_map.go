package main

import "fmt"

// 作为参数，是一个引用传递【传递的是一个地址，即作为一个指针】
func printMap(tMap map[string]string) {
	// 遍历
	for key, value := range tMap {
		fmt.Print("key = ", key, ",")
		fmt.Println("value = ", value)
	}
}

// 作为参数，是一个引用传递【传递的是一个地址，即作为一个指针】
func updateMap(tMap map[string]string) {
	tMap["USA"] = "DC"
}

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	printMap(cityMap)

	// 删除
	delete(cityMap, "Japan")

	// 修改
	updateMap(cityMap)

	printMap(cityMap)
}
