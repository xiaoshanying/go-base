//如何使用pool包来共享一组模拟的数据库连接
package main

import (
	"base01/pool/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	//使用的go routine数量
	maxGoRoutines = 25
	//池中的资源的数量
	pooledResources = 2
)

//模拟要共享的资源
type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close:Connection", dbConn.ID)
	return nil
}

//用来给每个连接分配一个独一无二的id
var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create:New Connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup

	wg.Add(maxGoRoutines)

	p, err := pool.New(createConnection, pooledResources)

	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoRoutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("shutdown Program.")

	p.Close()
}

//测试连接的资源池
func performQueries(query int, p *pool.Pool) {
	//从池里请求一个连接
	conn, err := p.Acquire()

	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	//模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
