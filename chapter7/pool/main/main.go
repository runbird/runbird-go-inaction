package main

import (
	"io"
	"log"
	"math/rand"
	"runbird-go-inaction/chapter7/pool"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutine   = 25
	pooledResource = 2
)

//模拟被共享的资源
type dbConnection struct {
	ID int32
}

func (conn *dbConnection) Close() error {
	log.Println("Close: Connection", conn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Accquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)

	//创建用来管理连接的池
	p, err := pool.New(createConnection, pooledResource)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()

	log.Println("Shutdown Program.")
	p.Close()
}
