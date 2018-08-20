package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

//读取并发序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	//打开文件
	file, err := os.Open(dataFile)

	if err != nil {
		return nil, err
	}

	//函数返回时执行(即使函数以外崩溃终止也会执行)
	defer file.Close()

	var feeds []*Feed

	//将文件解码到一个切片
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err

}
