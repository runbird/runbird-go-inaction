package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

//Run执行搜索
func Run(searchTerm string) {
	//获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	//创造无缓冲通道，接收匹配后的结果
	results := make(chan *Result)

	//构造waitgroup,用于多线程操作处理数据流
	var waitGroup sync.WaitGroup

	//设置需要等待处理
	//每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	//为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		//获取一个匹配器来查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		//启动一个goroutine来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	//启动一个goroutine来监控works 是否完成
	go func() {
		//等候所有任务完成
		waitGroup.Wait()

		//用关闭通道的方式，通知Display函数
		//退出程序
		close(results)
	}()

	//启动函数，显示返回结果 并且在最后一个结果显示完后返回
	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
