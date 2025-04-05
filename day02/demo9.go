package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}
func (u *User) Notify1() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}

func (u *User) Notify2() *User {

	return u
}
func (u User) Notify3() *User {

	return &u
}

func main() {
	// 值类型调用方法
	u1 := User{"golang", "golang@golang.com"}
	u1.Notify()
	u1.Notify1()
	// 指针类型调用方法
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()
	u3.Notify1()
}
