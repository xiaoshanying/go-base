package main

import (
	"base01/io/curl"
	"os"
)

/*
//创建一个Buffer值,并将一个字符串写入Buffer
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	//使用Fprintf来将一个字符串拼接到Buffer里
	//将bytes.Buffer的地址作为io.Writer类型值传入
	fmt.Fprintf(&b, "World!")

	//将Buffer的内容输出到标准输出设备
	b.WriteTo(os.Stdout)
*/
func main() {
	curl.Curl(os.Args[1], os.Args[2])
}
