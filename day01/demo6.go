package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Book1 struct {
	Name        string
	Author      string
	PublishTime time.Time
}

func main6() {
	bookObj := &Book1{
		Name:        "Go 语言基础入门",
		Author:      "fxtack",
		PublishTime: time.Now(),
	}

	// 注意 json.MarshalIndent 的参数多了两个
	jsonBytes, _ := json.MarshalIndent(bookObj, "", "\t")
	fmt.Println(string(jsonBytes))
}
