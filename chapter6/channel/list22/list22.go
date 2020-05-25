package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1
	wg.Wait()
}

//若Runner 前未加go，导致baton 无数据写入 runner，死锁
func Runner(baton chan int) {
	var newRunner int
	//等待接力棒 阻塞goroutine
	// 新的goroutine会被在这里阻塞，等待开启它goroutine传送值
	runner := <-baton
	fmt.Printf("Runner %d Running With Baton \n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To the Line \n", newRunner)
		//error Runner(baton)  becuase lose  go
		//baton 阻塞住gorotine
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished,Race Over \n", runner)
		wg.Done()
		return
	}

	//
	fmt.Printf("Runner %d Exchange With Runner %d \n", runner, newRunner)
	baton <- newRunner
}
