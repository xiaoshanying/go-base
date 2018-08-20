package main

import (
	_ "base01/sample/matchers"
	"base01/sample/search"
	"log"
	"os"
)

//init 函数在main函数之前执行
func init() {
	//日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
