package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Books struct {
	Name        string
	Author      string
	PublishTime time.Time
}

func main5() {
	bookObj := &Books{
		Name:        "Go 语言基础入门",
		Author:      "fxtack",
		PublishTime: time.Now(),
	}
	// 使用 encoding/json 包进行序列化
	// json.Marshal 序列化函数可以接受对象实例也可以接受对象指针
	// 返回 byte[] 与 error，使用 error 可以对序列化是否成功进行判断，此处省略
	jsonBytes, _ := json.Marshal(bookObj)
	// 以字符串形式将序列化结果输出
	fmt.Println(string(jsonBytes))
}
