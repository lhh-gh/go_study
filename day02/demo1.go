package main

import "fmt"

func main() {
	//初始化集合
	student := make(map[string]string)
	/* map插入key - value对 */
	student["name"] = "包子"
	student["age"] = "18"
	student["email"] = "baozi@163.com"
	fmt.Printf("len=%d type=%T ,map=%v\n", len(student), student, student)
	//直接赋值也是初始化，声明的时候填充元素
	var teacher = map[string]string{"name": "肉包子", "age": "20", "email": "roubaozi@163.com"}
	fmt.Printf("len=%d type=%T ,map=%v\n", len(teacher), teacher, teacher)
}
