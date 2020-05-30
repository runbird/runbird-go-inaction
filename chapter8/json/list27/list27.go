package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

//若返回的值为string形式存在，需要转换为byte[] ,使用json的 unmarshal 函数序列化
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c) //赋予c一个新的地址
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	//{Gopher programmer {415.333.3333 415.555.5555}}
	fmt.Println(c)
}
