package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	//使用buffer写入一个字符串
	var b bytes.Buffer
	//{[] 0 0}
	fmt.Println(b)

	b.Write([]byte("Hello "))

	//Fprintf 拼接
	fmt.Fprintf(&b, "World!")

	//os.File值的地址作为io.Writer类型值传入
	b.WriteTo(os.Stdout)
}
