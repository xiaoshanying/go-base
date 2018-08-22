package create

/*
基本案例
//分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Println("Start GoRoutines")

	go func() {
		//函数退出时调用done通知main函数工作已经完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		//函数退出时调用done通知main函数工作已经完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Waiting To finish")
	wg.Wait()
	fmt.Println("\nTerminal program")



	互吃所------

	var (
	counter int

	//同步
	wg sync.WaitGroup

	//互斥锁
	mutex sync.Mutex
)

func main() {

	wg.Add(2)

	go incCounter(1)

	go incCounter(2)

	wg.Wait()

	fmt.Printf("Final Counter: %d\n", counter)

}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			//强制切换go routine
			runtime.Gosched()

			value++

			counter = value
		}
		mutex.Unlock()
	}
}
*/
