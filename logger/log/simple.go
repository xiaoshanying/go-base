package log

/*

输出以下格式的日志
TRACE: 2009/11/10 23:00:00.000000 /tmpfs/gosandbox-/prog.go:14: message
//通常程序会在Init函数里配置日志参数
func init() {
	//日志项的前缀
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	//写到标准日志记录器
	log.Println("message")

	//Fatalln在调用println之后会接着调用os.Exit(1)
	log.Fatalln("fatal message")

	//在调用println之后会接着调用panicl
	log.Panicln("panic message")
}

*/
