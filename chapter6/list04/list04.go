package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

//逻辑处理器会避免某一个goroutine长时间占用逻辑处理器（逻辑处理器调用操作系统的线程执行）
//A:3769
//A:3779
//A:3793
//B:3691
//B:3697
//可以看出有时间片的切换
func main() {
	runtime.GOMAXPROCS(1) // 切换的较少
	//runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	fmt.Println("create goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("waiting to finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

//显示5000以内的素数
func printPrime(prefix string) {
	defer wg.Done()
	//next:
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
