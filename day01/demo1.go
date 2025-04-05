package main

import (
	"fmt"
	"time"
)

func main1() {
	now := time.Now()
	formatted := now.Format("2006-01-02 15:04:05 MST")
	fmt.Println(formatted)
}
