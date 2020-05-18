package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//更改调度器设定一个逻辑处理器
	runtime.GOMAXPROCS(1)

	//设定2个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start Goroutines")

	//声明一个匿名函数，并且创建一个goroutine
	go func() {
		//函数退出时，通知main函数进行 wg -1
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nterminalting Program")
}
