package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("eat ? ")
}

func (this *Human) Walk() {
	fmt.Println("walk ")
}

// 结构体的嵌套(可以类似java的对象继承，但是只是形式上可以类似)
type SuperMan struct {
	human Human
	level int
}

// 形式上可以看成重写，实际只是一个不同的函数
func (this *SuperMan) Eat() {
	this.human.Eat()
	fmt.Println("food")
}

// 形式上可以看成子类的新方法，实际只是一个不同的函数
func (this *SuperMan) Fly() {
	fmt.Println("fly ... ")
}

func main() {

	person := Human{sex: "男", name: "666"}
	person.Eat()
	person.Walk()

	superMan := SuperMan{human: Human{sex: "？", name: "6"}, level: 12}
	superMan.Eat()
	superMan.human.Walk()
	superMan.Fly()
}
