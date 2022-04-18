package main

func main () {
	// 常量(只读)
	const len int = 10
	// 用常量定义枚举类型
	const (
		// 可以在const()里添加关键字iota,每一行iota会累加1，第一行默认是0
		BEIJING = iota
		SHENZHEN
		GUANGZHOU
	)
	// iota定制公式
	const(
		a, b = iota + 1, iota + 2  // iota = 0, a=0+1,b=0+2
		c, d                       // iota = 1, c=1+1,d=1+2
		e, f                       // iota = 2, e=2+1,f=2+2
		// 改变了递增公式
		g, h = iota * 2, iota * 3  // iota = 3, a=3*2,b=3*3
		i, k                       // iota = 4, a=4*2,b=4*3
	)
}