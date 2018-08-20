package search

import (
	"fmt"
	"log"
)

//保存搜索的结果
type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)

	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

//display 从每个单独的go routine 接收到结果后，在终端窗口输出
func Display(results chan *Result) {
	//通道会一直被阻塞,直到有结果写入
	//一旦通道被关闭,for循环就会终止
	for result := range results {
		fmt.Println("%s:\n%s\n\n", result.Field, result.Content)
	}
}
