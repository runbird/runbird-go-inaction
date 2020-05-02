package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(contents chan int)  {
	for content := range contents{
		fmt.Printf("Received %d ", content)
	}
	wg.Done()
}

func main()  {
	contents:= make(chan int)
	go printer(contents)
	wg.Add(1)

	for i := 1; i <= 10; i++ {
		contents <- i;
	}
	close(contents)
	wg.Wait()
}
