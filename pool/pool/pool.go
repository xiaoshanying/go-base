//管理用户定义的一组资源
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

/*
管理一组可以安全的在多个go routine间共享的资源
被管理的资源必须实现io.Closer接口
*/
type Pool struct {
	//互斥锁保证多个go routine访问资源池时,池内的值是安全的
	m sync.Mutex

	//通道类型可以是接口,池可以管理任意实现了io.Close的资源
	resources chan io.Closer

	//需要由包的使用者提供
	factory func() (io.Closer, error)

	//标识池是否被关闭
	closed bool
}

//请求了一个已经关闭的池时返回此错误
var ErrPoolClosed = errors.New("Pool has been closed.")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

//从池中获取一个资源,有空闲资源返回,没有则创建
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	//检查通道是否有可用资源
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource.")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resource.")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	//如果池已经被关闭,销毁这个资源
	if p.closed {
		r.Close()
		return
	}

	//试图将释放的资源放回池中,池子满了则关闭这个资源
	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue.")
	default:
		log.Println("Release:", "Closing.")
		r.Close()
	}
}

//让资源池停止工作,并关闭所有现有的资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	//如果已经关闭,则什么也不做
	if p.closed {
		return
	}

	//关闭
	p.closed = true

	//在清空通道里的资源之前,将通道关闭,如果不这样做,会发生死锁
	close(p.resources)

	//关闭资源
	for r := range p.resources {
		r.Close()
	}
}
