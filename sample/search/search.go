package search

import (
	"log"
	"sync"
)

//注册用于搜索的匹配器映射
var matchers = make(map[string]Matcher)

//执行搜索逻辑
func Run(searchTerm string) {
	//获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	//创建一个无缓冲的通道,接受匹配后的结果
	results := make(chan *Result)

	//构造一个waitGroup,以便处理所有的数据源
	var waitGroup sync.WaitGroup

	//设置需要等待处理的每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		//启动一个go routine来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	//启动一个go routine来监控是否所有的工作都做完了
	go func() {
		//等待所有任务完成,会阻塞go routine,直到waitGroup内部计数达到0
		waitGroup.Wait()

		//用关闭通道的方式,通知Display函数,可以退出程序了
		close(results)
	}()

	//启动函数,显示返回的结果,并且在最后一个结果显示完后返回
	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
