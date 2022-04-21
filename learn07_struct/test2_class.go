package main

import "fmt"

type Hero struct {
	// 姓名
	name string
	// 攻击力、等级
	ad, level int
}

func (this *Hero) GetName() {
	fmt.Println("name = ", this.name)
}

func (this *Hero) SetName(name string) {
	this.name = name
}

func main() {
	// 创建一个结构体
	hero := Hero{name: "666", ad: 20, level: 1}
	// 获取姓名
	hero.GetName()
	// 修改姓名
	hero.SetName("777")
	hero.GetName()
}
