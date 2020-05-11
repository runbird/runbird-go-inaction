package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	//将字符串写入Buffer
	b.Write([]byte("Hello"))

	//使用Fprintf将字符串拼接到Buffer
	fmt.Fprintf(&b, " World!")

	io.Copy(os.Stdout, &b)
}
