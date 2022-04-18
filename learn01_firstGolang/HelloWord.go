package main   // 当前程序所在的包名

// 导入包，导入多个也可以选择分开写
import (
	"fmt"
	"time"
)


// main函数，golang强制性要求函数的{一定和函数名在同一行，否则编译错误
func main() {
	// golang中表达式可加;也可以不加
	fmt.Println("hello GO!")
	// 休眠一秒
	time.Sleep(1 * time.Second)
}