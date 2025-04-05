package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Bookss struct {
	// 使用 json 标签
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	PublishTime time.Time `json:"publishTime"`
}

func main() {
	bookObj := &Bookss{
		Name:        "Go 语言基础入门",
		Author:      "fxtack",
		PublishTime: time.Now(),
	}

	// 调用序列化函数，序列化结果存储到 ./book.json
	WriteJson("./book.json", bookObj)
}

// 序列化为 .json 文件函数
// filePath 为存储到的路径，ref 为需要序列化的对象或其应用
func WriteJson(filePath string, ref interface{}) {

	// 判断提供的路径文件是否存在，避免直接覆盖已有的文件
	f, err := os.Stat(filePath)
	if (err == nil || os.IsExist(err)) && !f.IsDir() {
		log.Println("文件已经存在或存在同名文件夹")
		return
	}

	// 创建 .json 文件
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("文件创建失败")
	}
	defer file.Close()

	// 序列化
	json, err := json.MarshalIndent(ref, "", "\t")
	if err != nil {
		log.Println("序列化失败")
	}

	// 写入 .json 文件
	file.Write(json)
}
