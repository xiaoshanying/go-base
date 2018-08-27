package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger //记录所有日志
	Info    *log.Logger //记录重要的信息
	Warning *log.Logger //需要注意的信息
	Error   *log.Logger //非常严重的问题
)

//通常程序会在Init函数里配置日志参数
func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	//New创建并正确初始化一个Logger类型的值,函数New会返回新创建的值的地址
	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	Trace.Println("I have sm st to say")

	Info.Println("Special Information")

	Warning.Println("There is a warning")

	Error.Println("something has failed")

}
