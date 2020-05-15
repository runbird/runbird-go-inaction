package main

import (
	"fmt"
	"runbird-go-inaction/chapter5/list68/counters"
)

func main() {
	//短变量声明操作符 := 创建未公开标识符
	counter := counters.New(10)
	fmt.Printf("Counter: %d\n", counter)
}
