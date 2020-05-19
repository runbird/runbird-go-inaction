package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

//使用atomic包中的 storeInt64 和 LoadInt64来同步
func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	//主线程等待两goroutine执行1s
	time.Sleep(1 * time.Second)

	//该工作停止，安全地设置了shutdown标志
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}

//dowork模拟执行工作的goroutine
//检测之前的shutdown标志来决定是否提前终止
func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Println("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		//要停止工作了吗
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
