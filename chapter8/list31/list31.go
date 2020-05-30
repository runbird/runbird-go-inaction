package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}
	//带缩进的形式
	indent, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(indent))

	//不带缩进
	result, err := json.Marshal(c)
	if err != nil {
		return
	}
	fmt.Println(string(result))
}
