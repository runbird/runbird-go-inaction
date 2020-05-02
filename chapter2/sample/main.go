package main

import (
	"log"
	"os"
	//github这部分可替换为runbird-go-inaction
	//_ "github.com/goinaction/code/chapter2/sample/matchers"
	//"github.com/goinaction/code/chapter2/sample/search"
	"runbird-go-inaction/chapter2/sample/search"
	_ "runbird-go-inaction/chapter2/sample/matchers"  //使用别名调用 matchers/rss.go 中init()方法
)

//init在main之前调用
func init() {
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

//整个程序的入口
func main() {
	//使用特定的项做搜索
	search.Run("president")
}
