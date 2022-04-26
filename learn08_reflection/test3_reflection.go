package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (receiver User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", receiver)
}

func (receiver User) Add(a, b int) (c, d int) {
	return a + b, b - a
}

func main() {
	user := User{1, "666", 18}
	user.Call()
	fmt.Println("-------------------------------")
	GetFieldAndMethod(user)
}

func GetFieldAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is : ", inputType)
	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value is : ", inputValue)
	fmt.Println("-------获取input的内容字段-------")
	// 1.获取input的type，通过NumField，获取字段数进行遍历
	// 2.获得每个字段field 的数据类型
	// 3.通过field的一个Interface()方法获得对应value值
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)  // 类型为 var field StructField
		value := inputValue.Field(i) // 类型为 var value Value
		// 只能访问公共属性，访问私有会报错，需要设置安全检查值为false[做一些额外配置]
		// field.Name 获取字段的名称
		// value.Interface() 获取字段的值
		fmt.Printf("%s : %v\n", field.Name, value.Interface())
	}
	fmt.Println("-----获取input里面的方法以及调用-----")
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i) // 类型为 var method Method
		// 函数名称，函数签名
		methodName := method.Name // 类型为 var methodName string
		methodType := method.Type // 类型为 var methodType Type
		fmt.Printf("%s: %v\n", methodName, methodType)
		if strings.EqualFold(method.Name, "Call") {
			// 无参数的调用方法
			methodValue := inputValue.MethodByName(methodName)
			methodValue.Call(nil)
		} else {
			// 有参数的调用方法
			methodValue := inputValue.MethodByName(methodName)
			params := []reflect.Value{
				reflect.ValueOf(1),
				reflect.ValueOf(5),
			}
			methodValue.Call(params)
		}
	}
}
