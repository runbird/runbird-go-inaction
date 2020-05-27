package main

import (
	"log"
	"os"
	"runbird-go-inaction/chapter7/runner"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")
	//分配任务执行时间
	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeOut:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interruput.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

//内部的i 刚开始没有传值？
func createTask() func(int) {
	return func(i int) {
		log.Printf("Processor - Task #%d.", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}
