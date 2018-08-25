//管理一个go routine池来完成工作
package work

import "sync"

type Worker interface {
	Task()
}

//go routine池,这个池可以完成任何已提交的worker任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

//创建一个新工作池
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			//一旦通道关闭,循环就会结束
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

//将工作提交到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

//shuodown等待所有的go routine停止工作
func (p *Pool) Shutdown() {
	//关闭通道,池里的所有go routine会停止工作
	close(p.work)
	p.wg.Wait()
}
