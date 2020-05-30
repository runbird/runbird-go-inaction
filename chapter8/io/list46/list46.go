package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	//创建文件来保存响应内容
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	//MultiWriter,同时向文件和标准输出设备进行写操作
	dest := io.MultiWriter(os.Stdout, file)

	//读出响应内容并 写到两个目的地
	io.Copy(dest, resp.Body)
	if err := resp.Body.Close(); err != nil {
		log.Println(err)
	}
}
