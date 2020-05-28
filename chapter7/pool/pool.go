package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed.")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}
	return &Pool{
		//分配池资源大小
		resources: make(chan io.Closer, size),
		factory:   fn,
	}, nil
}

func (p *Pool) Accquire() (io.Closer, error) {
	select {
	//检查是否有空闲资源
	case r, ok := <-p.resources:
		log.Println("Accquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		//没有空闲的资源可用，提供新的资源
		log.Println("Accquire:", "New Resource")
		//p.factory =  func() (io.Closer, error)   p.factory() = (io.Closer, error)
		return p.factory()
	}
}

//将使用后的资源放入池内
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和Close操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	//如果池已经关闭，则销毁该资源
	if p.closed {
		r.Close()
		return
	}

	select {
	//试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	//如果队列已经满了，则关闭这个资源
	default:
		log.Println("Releas:", "Closing")
		r.Close()
	}
}

//Close让资源池停止工作，并且关闭现有的所有资源
func (p *Pool) Close() {
	//保证本操作与Release操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	//如果pool已经被关闭，什么也不用做
	if p.closed {
		return
	}

	//将池关闭
	p.closed = true
	//清空通道里的资源之前，将通道关闭，否则会发生死锁
	close(p.resources)

	//关闭资源
	for r := range p.resources {
		r.Close()
	}
}
