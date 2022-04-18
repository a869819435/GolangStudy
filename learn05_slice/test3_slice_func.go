package main

import "fmt"

func main()  {
	numbers := make([]int, 3 ,5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers1 := make([]int, 3)
	// 默认的cap会与len相等
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers1), cap(numbers1), numbers1)

	// 追加
	numbers = append(numbers, 1,2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	// 当追加容量超过cap时，自动扩容，扩容的大小为cap的大小
	numbers = append(numbers, 1,2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	s := []int{1,2,3,4,5}
	// 截取  [startIndex,endIndex) 之间的元素
	s1 := s[0:2]
	fmt.Println(s1)
	s2 := s[2:4]
	fmt.Println(s2)

	s[0] = 5
	fmt.Println(s1)
	s3 := make([]int, 5)
	// 深拷贝
	copy(s3, s)
	s3[0] = 100
	fmt.Println(s3)
}
