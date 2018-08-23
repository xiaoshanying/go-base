package run

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//在给定的时间内执行一组任务,并在操作系统发送中断信号时结束这些任务
type Runner struct {
	//从操作系统发出的信号
	interrupt chan os.Signal
	//报告任务是否完成
	complete chan error
	//任务超时
	timeout <-chan time.Time
	//存储任务
	tasks []func(int)
}

//任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

//收到操作系统事件时返回
var ErrInterrupt = errors.New("received interrupt")

//创建一个新的runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//添加一组任务到runner
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//开始执行runner
func (r *Runner) Start() error {

	//接收操作系统的信息号,并发送到通道
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	//检测到任务完成,或者超时
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

//任务执行
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//如果接收到中断信号,则中断
		if r.goInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) goInterrupt() bool {
	//判断是否接收到了中断信号
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
