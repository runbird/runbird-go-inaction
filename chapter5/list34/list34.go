package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	//Body 实现了io.Reader接口； 传入io.Copy(Writer,Reader)
	//Body 是io.ReadCloser类型，该类型是Reader 、 Closer接口
	io.Copy(os.Stdout, resp.Body)
	if err := resp.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
