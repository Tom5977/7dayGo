package main

import (
	"fmt"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func main() {
	u1 := User{}
	var in interface{} = u1
	fmt.Printf("%T", in)
}
