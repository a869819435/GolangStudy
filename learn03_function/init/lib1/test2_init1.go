package lib1

import "fmt"

// 开头大写表示公有函数
func LibFunc1()  {
	fmt.Println("lib1.libFunc1 ... OK!")
}

func init()  {
	fmt.Println("start init1 ...")
}
