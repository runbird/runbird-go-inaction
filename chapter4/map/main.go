package main

import (
	"fmt"
)

//映射
func main() {
	//key-string / value-int
	dict := make(map[string]int)
	dict2 := map[string]string{"Red": "#da1337", "Orage": "#e95a22"}
	//声明一个 存储字符串切片的映射
	dict3 := map[int][]string{}
	fmt.Println(dict, "\n", dict2, "\n", dict3)

	colors := map[string]string{}
	colors["Red"] = "#da1337"

	//错误操作 nil
	var colorsError map[string]string
	colorsError["Red"] = "#da1337"

	//判断是否存在 及 值
	value, exists := colors["Blue"]
	if exists {
		fmt.Println(value)
	}

	//从映射获取值，通过值判断键是否存在
	value2 := colors["Blue"]
	if value2 != "" {
		fmt.Println(value2)
	}

	//函数间传递
	colorsRange := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "ff7f50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	for key, value := range colorsRange {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
	removeColor(colorsRange, "Coral")
	//未使用副本，使用小数据量进行修改 类似 slice
	for key, value := range colorsRange {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
