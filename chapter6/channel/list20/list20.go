package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//无缓冲的channel  同步、阻塞
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//内置函数make(chan关键字   共享的数据类型   是否有缓冲)
	//可共享 内置类型 命名类型 结构类型 和 引用类型的值或者指针
	court := make(chan int)
	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	//发球
	court <- 1
	wg.Wait()
}

//player模拟一个选手在打网球
func player(name string, court chan int) {
	defer wg.Done()

	for {
		//等待球被打过来
		ball, ok := <-court
		if !ok {
			//通道被关闭，赢球
			fmt.Printf("Player %s Won \n", name)
		}

		//选择随机数来判定是否丢球
		n := rand.Intn(100)
		if n%3 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			//关闭通道，表示我们输了
			close(court)
			return
		}

		//显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		//将球打向对手
		court <- ball
	}
}
