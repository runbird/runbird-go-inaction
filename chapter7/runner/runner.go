package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	//通道报告从操作系统发送的信号
	interrupt chan os.Signal

	//通道报告处理任务已经完成
	complete chan error

	//报告处理任务已经超时
	timeOut <-chan time.Time
	//func切片
	tasks []func(int)
}

var ErrTimeOut = errors.New("received timeout")
var ErrInterrupt = errors.New("received Interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		//带缓冲的通道  保证通道至少能接收一个来自语言运行时的os.Signal值
		//确保运行时发送这个事件不会被阻塞
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		// time.After()   返回的是  <-chan time.Time
		timeOut: time.After(d),
		//task 初始值设置为nil
	}
}

//将任务通过Add 加到Runner上
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//Start执行所有的任务，并监视事件
func (r *Runner) Start() error {
	//接收所有的中断信号
	signal.Notify(r.interrupt, os.Interrupt)
	//用不通的goroutine执行不通的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeOut:
		return ErrTimeOut
	}
}

//run执行一个已经注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//检测系统的中断信号
		if r.gotIntterupt() {
			return ErrInterrupt
		}
		//执行已经注册的id
		task(id)
	}
	return nil
}

func (r *Runner) gotIntterupt() bool {
	select {
	//当中断时间被触发时发出的信号
	case <-r.interrupt:
		//停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
		//继续正常运行
	default:
		return false
	}
}
