package search

import (
	"log"
)

//Result保存搜索的结果
type Result struct {
	Field   string
	Content string
}

//Matcher定义了要实现的
//新搜索类型的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//goroutine并发执行
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//对特定的匹配器执行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	//将结果写入通道
	for _, result := range searchResults {
		results <- result
	}
}

// Display从每个单独的goroutine接收到结果后 终端输出
func Display(results chan *Result) {
	//通道会一直阻塞至结果写入
	//通道关闭的同时 close(results)，for循环停止
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
