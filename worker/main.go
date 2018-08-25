package main

import (
	"base01/worker/work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"jaos",
	"jason",
}

type namePrinter struct {
	name string
}

//实现worker接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	//创建工作池
	p := work.New(2)

	var wg sync.WaitGroup

	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}
