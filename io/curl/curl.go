package curl

import (
	"io"
	"log"
	"net/http"
	"os"
)

//curl get获取内容
func Curl(url string, destFile string) (code int) {
	r, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return r.StatusCode
	}

	file, err := os.Create(destFile)
	if err != nil {
		log.Println(err)
		return 1
	}

	defer file.Close()

	//使用MultiWriter可以同时向文件和标准输出设备进行写操作
	dest := io.MultiWriter(os.Stdout, file)

	//读出响应内容,并写到两个目的地
	io.Copy(dest, r.Body)

	if err := r.Body.Close(); err != nil {
		log.Println(err)
		return 2
	}
	return 0
}
