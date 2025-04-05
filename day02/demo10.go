package main

import "fmt"

type User1 struct {
	id   int
	name string
}
type Manager struct {
	User1
}

func (self *User1) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}
func main() {
	m := Manager{User1{1, "zhangsan"}}
	fmt.Println("Manager:", &m)
	fmt.Println(m.ToString())

}
