package main

import (
	"base01/runner/run"
	"log"
	"os"
	"time"
)

//规定任务必须在3秒内完成
const timeout = 3 * time.Second

func main() {

	log.Println("Starting work..")

	r := run.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case run.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case run.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")

}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
