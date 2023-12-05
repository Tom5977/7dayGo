package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func main() {
	u1 := User{}
	u1ptr := &User{}
	var uu []User
	// fmt.Println(reflect.Indirect(reflect.ValueOf(&uu)).Type().Elem(), reflect.Indirect(reflect.ValueOf(u1ptr)).Type())
	fmt.Println(reflect.TypeOf(u1), reflect.TypeOf(u1ptr))
	fmt.Printf("%T\n", reflect.Indirect(reflect.New(reflect.ValueOf(u1).Type().Field(1).Type)))
	fmt.Println(reflect.ValueOf(uu).Type().Elem())
	dT := reflect.ValueOf(uu).Type().Elem()
	inF := reflect.New(dT).Elem().Interface()
	fmt.Println(reflect.TypeOf(inF))
	fmt.Println(reflect.ValueOf(uu).Type() == reflect.TypeOf(uu))

	fmt.Println(reflect.Indirect(reflect.ValueOf(&uu)))
}
