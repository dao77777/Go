package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (this *User) Call() {
	fmt.Println("Call()...")
	fmt.Println(this.Id)
}

type Resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex string `info:"sex"`
}

func findTag(resume *Resume) {
	t := reflect.TypeOf(resume).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagString := t.Field(i).Tag.Get("info")
		fmt.Println("info:", tagString)
	}
}

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

func DoFiledAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n",field.Name, field.Type, value)
	}
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func main() {
	reflectNum(3.144)
	fmt.Println("------------")
	reflectNum(3)
	fmt.Println("------------")

	user := User{0, "dao77777", 26}
	DoFiledAndMethod(user)
	user.Call()

	findTag(&Resume{Name: "dao77777", Sex: "male"})
}