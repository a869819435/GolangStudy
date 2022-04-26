package main

import "fmt"

// interface本质是一种指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的结构体
type Cat struct {
	Color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat sleep")
}

func (this *Cat) GetColor() string {
	return this.Color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的结构体
type Dog struct {
	Color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dag sleep")
}

func (this *Dog) GetColor() string {
	return this.Color
}

func (this *Dog) GetType() string {
	return "Dag"
}

func show(animalIF AnimalIF) {
	animalIF.Sleep()
	fmt.Println("color = ", animalIF.GetColor())
	fmt.Println("kind = ", animalIF.GetType())
}

func main() {
	var animal AnimalIF // 接口的数据类型
	animal = &Cat{Color: "Blue"}
	animal.Sleep()

	cat := Cat{"Blue"}
	dog := Dog{"Red"}
	show(&cat)
	show(&dog)
}
