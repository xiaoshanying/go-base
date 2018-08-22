package channel

/*
当一个资源在go routine之间共享时,通道在go routine之间架起了一个管道
并提供了确保同步交换数据的机制

1.创建一个无缓冲的通道
buffered := make(chan int)

2.创建一个有缓冲的通道
buffered := make(chan string,10)

向通道发送一个字符串
buffered <- "hello"

从通道接收数据
value := <- buffered


无缓冲的通道模拟打网球
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court)

	go player("Mike", court)

	//发球
	court <- 1

	wg.Wait()

}

//模拟打网球
func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court

		if !ok {
			fmt.Printf("player %s Won\n", name)
			return
		}

		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("player %s Missed\n", name)
			close(court)
			return
		}

		fmt.Printf("player %s hit %d\n", name, ball)
		ball++

		court <- ball
	}
}


无缓冲通道模拟跑步
var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("Runner %d Running with Baton\n", runner)

	//创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To the Line\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished,Race Over\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner

}

*/
