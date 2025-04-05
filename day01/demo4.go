package main

import (
	"fmt"
	"time"
)

type Book struct {
	Name        string
	Author      string
	PublishTime time.Time
}

func main4() {
	bookObj := &Book{
		Name:        "Go 语言基础入门",
		Author:      "fxtack",
		PublishTime: time.Now(),
	}
	fmt.Println(bookObj)
}
