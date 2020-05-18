package main

import (
	"fmt"
	"runtime"
	"sync"
)

//竞争状态 race candition
// 变量counter会进行4次读写操作，每个goroutine会执行两次
var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		//当前goroutine从线程退出，并放回队列
		runtime.Gosched()
		value++
		counter = value
	}
}
