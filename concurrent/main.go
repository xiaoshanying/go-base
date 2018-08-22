package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	//使用 go routine的数量
	numberGoroutines = 4
	//任务数
	taskLoad = 10
)

var wg sync.WaitGroup

func init() {
	//初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {
	//创建有缓冲的通道
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	//启动go routine处理工作
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	//所有工作处理完,关闭通道
	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks

		if !ok {
			fmt.Printf("Worker: %d :Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker %d :Started %s\n", worker, task)

		sleep := rand.Int63n(100)

		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker : %d :Completed %s\n", worker, task)
	}
}
