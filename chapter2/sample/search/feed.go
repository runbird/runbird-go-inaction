package search

import (
	"encoding/json"
	"os"
)

const dataFile = "chapter2/sample/data/data.json"

// `` --> tag;标记里描述了JSON解码的元数据
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	//打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	//当函数返回时
	//关闭文件
	defer file.Close()

	//将文件解码到一个切片里
	//切片中的每一项都是指向一个Feed类型的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	//这个函数不需要检查错误，调用者会做这件事
	return feeds, err
}
