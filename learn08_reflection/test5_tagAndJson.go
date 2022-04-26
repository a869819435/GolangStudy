package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Price  int    `json:"price"`
	Actors []string
}

func main() {
	movie := Movie{"喜剧之王", 2000, 20, []string{"周星驰", "张柏芝"}}
	// 编码过程  结构体转为json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error,reason: ", err)
		return
	}
	fmt.Printf("json = %s\n", jsonStr)
	// 编码过程  json转为结构体
	my := Movie{}
	err = json.Unmarshal(jsonStr, &my)
	if err != nil {
		fmt.Println("json unmarshal error ", err)
		return
	}
	fmt.Printf("%v\n", my)
}
