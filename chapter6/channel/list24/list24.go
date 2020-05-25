package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用有缓冲的chan
const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	//有缓冲的channel
	task := make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(task, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		//task <- fmt.Printf("Task :%d", post) complie error: 需要转化为string 传输
		task <- fmt.Sprintf("Task :%d", post)
	}

	//仍然可以从关闭的通道中接收数据，但是不能再发送数据。
	//从一个已关闭且没有数据的通道获取数据，总会立刻返回，并返回一个通道类型的零值
	close(task)
	wg.Wait()
}

//worker作为goroutine 启动来处理 从有缓冲的通道传入的工作
func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//通道已经为空，并且已经被关闭
			fmt.Printf("Worker: %d closed\n", worker)
			return
		}

		fmt.Printf("Worker: %d Started  %s\n", worker, task)

		sleep := rand.Intn(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d Completed:%s\n", worker, task)
	}
}
